package api

import (
	"backend/comm"
	"backend/database"
	finance "backend/financizer"
	"backend/html_report"
	"backend/parser"
	helpers "backend/rabbitmq"
	"backend/search"
	"encoding/json"

	"net/http"
	"os"
	"strconv"
	"strings"

	// "encoding/base64"
	"io"
	"path"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type FileAndMetric struct {
	File   string
	Metric string
}

// func temp() {

// 	currentWorkingDirectory, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	file, err := os.Open(path.Join(currentWorkingDirectory, "bio.htm"))

// 	bytes, err := io.ReadAll(file)

// 	processReport(&FileAndMetric{
// 		File:   string(bytes),
// 		Metric: "Status",
// 	})
// }

func InitAPI() {
	// temp()

	createApi()
}

func processReport(fileAndMetric *FileAndMetric) uint {
	t := parser.InitParser([]byte(fileAndMetric.File))
	m := finance.InitFinancer(t, fileAndMetric.Metric)
	return html_report.InitHtmlReport(m, []byte(fileAndMetric.File))
}

const API = "api"

func createApi() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// Access-Control-Allow-Origin
		AllowMethods:     []string{"PUT", "GET", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
		AllowFiles:       true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST(path.Join(API, "post", "document", "add-metrics"), func(c *gin.Context) {
		bytes, err := io.ReadAll(c.Request.Body)
		helpers.FailOnError(err, "failed to parse")

		fileAndMetric := &FileAndMetric{}
		err = json.Unmarshal(bytes, fileAndMetric)

		if err != nil || fileAndMetric.File == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": "failed to add metrics to document",
			})
			return
		}

		// get document
		document := &comm.Document{}
		database.Db.Model(comm.Document{}).Preload("Metrics").Find(document, fileAndMetric.File)
		// get metrics
		t := parser.InitParser(document.Report)
		moreMetrics := finance.InitFinancer(t, fileAndMetric.Metric)
		// add metrics to document and save it
		database.CloneMetrics(moreMetrics)
		document.Metrics = append(document.Metrics, moreMetrics...)
		database.Db.Save(document)

		c.JSON(http.StatusOK, gin.H{
			"id": document.ID,
		})

		log.Info("Document id sent")
	})

	r.POST(path.Join(API, "post", "document"), func(c *gin.Context) {
		bytes, err := io.ReadAll(c.Request.Body)
		helpers.FailOnError(err, "failed to parse")

		fileAndMetric := &FileAndMetric{}
		err = json.Unmarshal(bytes, fileAndMetric)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to parse",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"id": processReport(fileAndMetric),
		})
		log.Info("Document id sent")
	})

	type NewDoc struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	r.POST(path.Join(API, "post", "document", "change-name"), func(c *gin.Context) {
		bytes, err := io.ReadAll(c.Request.Body)
		helpers.FailOnError(err, "failed to parse")

		newDoc := &NewDoc{}
		err = json.Unmarshal(bytes, newDoc)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "did not send id",
			})
		}

		document := &comm.Document{}
		database.Db.Model(comm.Document{}).Find(document, newDoc.ID)
		document.Name = newDoc.Name
		database.Db.Save(document)

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	type Id struct {
		ID string `json:"id"`
	}

	r.POST(path.Join(API, "post", "document", "get-name"), func(c *gin.Context) {
		bytes, err := io.ReadAll(c.Request.Body)
		helpers.FailOnError(err, "failed to parse")

		id := &Id{}
		err = json.Unmarshal(bytes, id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "did not send id",
			})
		}

		document := &comm.Document{}
		database.Db.Model(comm.Document{}).Find(document, id.ID)

		c.JSON(http.StatusOK, gin.H{
			"message": document.Name,
		})
	})

	r.GET(path.Join(API, "get", "document/", ":id"), func(c *gin.Context) {
		// retrieve document with id
		id := strings.ReplaceAll(c.Param("id"), "\"", "")
		nextId := strings.ReplaceAll(id, "$", "")

		document := &comm.Document{}
		database.Db.Model(comm.Document{}).Find(document, nextId)

		c.Data(http.StatusOK, "text/html; charset=utf-8", document.Report)
	})

	r.GET(path.Join(API, "get", "document/", "all"), func(c *gin.Context) {
		// retrieve all document ID's

		documents := []*comm.Document{}

		database.Db.Model(&comm.Document{}).Find(&documents).Select("ID")

		ids := []uint{}
		names := []string{}

		for _, doc := range documents {

			ids = append(ids, doc.ID)
			names = append(names, doc.Name)
		}

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

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

		u64, err := strconv.ParseUint(c.Param("docid"), 10, 64)
		if metric.DocumentID == 0 {
			metric.DocumentID = uint(u64)
		}

		// get old metrics
		document := comm.Document{}
		database.Db.Model(&comm.Document{}).
			Preload("Metrics").
			Preload("Metrics.Submetric").
			First(&document, c.Param("docid"))

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

		document := comm.Document{}

		database.Db.Model(&comm.Document{}).
			Preload("Metrics").
			Preload("Metrics.Submetric").
			First(&document, c.Param("docid"))

		comm.OrderSubmetrics(document.Metrics)

		c.JSON(http.StatusOK, gin.H{
			"message": document.Metrics,
		})

	})

	r.GET(path.Join("assets", "NeueMachina-Regular-e896c98c.otf"), func(c *gin.Context) {

		workingDirectory, err := os.Getwd()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad",
			})
			return
		}

		path := path.Join(workingDirectory, "html_report", "ui", "dist", "fonts", "NeueMachina-Regular.otf")
		bytes, err := os.ReadFile(path)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad",
			})
			return
		}

		log.Info("Sent Font")

		c.Data(http.StatusOK, "application/font-sfnt; charset=utf-8", bytes)

	})

	r.GET(path.Join("search", "get"), func(c *gin.Context) {

		// var buffer *bytes.Buffer

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		path := path.Join(dir, "search.html")
		file, err := os.Create(path)
		defer file.Close()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad",
			})
		}

		search.InitSearch().Component.Render(c, file)

		bytes, err := os.ReadFile(path)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad",
			})
		}

		c.Data(http.StatusOK, "text/html", bytes)

	})

	r.Run(":8080")
}
