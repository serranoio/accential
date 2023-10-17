package api

import (
	"backend/comm"
	"backend/database"
	"encoding/json"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func metricRoutes(r *gin.Engine) {
	r.GET(path.Join(API, "get", "metrics", "all"), func(c *gin.Context) {
		var allTags []comm.Tag

		database.Db.Find(&allTags)

		ids := []uint{}
		names := []string{}

		for _, tag := range allTags {
			foundName := false
			for _, name := range names {
				if tag.Name == name {
					foundName = true
				}
			}
			if !foundName {
				ids = append(ids, tag.ID)
				names = append(names, tag.Name)

			}

		}

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})

	})

	// this is currently for ALL documents. Make it for specific document
	r.POST(path.Join(API, "metric", "add", ":docid"), func(c *gin.Context) {

		id := strings.ReplaceAll(c.Param("docid"), "$", "")
		bytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"failed to parse": "bad",
			})
		}

		// get metric & update it with ID
		metric := &comm.Metric{}
		err = json.Unmarshal(bytes, &metric)

		var metrics []*comm.Metric
		metrics = append(metrics, metric)

		u64, err := strconv.ParseUint(id, 10, 64)
		if metric.DocumentID == 0 {
			metric.DocumentID = uint(u64)
		}

		// get old metrics
		document := comm.Document{}
		database.Db.Model(&comm.Document{}).
			Preload("Metrics").
			Preload("Metrics.Submetric").
			First(&document, id)

		// update old one if theyre same ID
		for i, documentMetrics := range document.Metrics {
			if documentMetrics.ID == metric.ID {
				document.Metrics[i] = metric
				document.Metrics[i].Submetric = metric.Submetric
			}
		}

		if metric.DocumentID == 0 {
			// combine them
			document.Metrics = append(document.Metrics, metrics...)
		}

		// save it all
		database.Db.Save(document)
		database.Db.Save(metric)
		database.Db.Save(metric.Submetric)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"failed to parse": "bad",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.GET(path.Join(API, "metric", "get", ":docid"), func(c *gin.Context) {

		id := strings.ReplaceAll(c.Param("docid"), "$", "")

		document := comm.Document{}

		database.Db.Model(&comm.Document{}).
			Preload("Metrics").
			Preload("Metrics.Submetric").
			First(&document, id)

		comm.OrderSubmetrics(document.Metrics)

		log.Info("Sent all metrics", document.Metrics)

		c.JSON(http.StatusOK, gin.H{
			"message": document.Metrics,
		})

	})

}
