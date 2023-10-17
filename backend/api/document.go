package api

import (
	"backend/comm"
	"backend/database"
	finance "backend/financizer"
	"backend/parser"
	"encoding/json"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func documentRoutes(r *gin.Engine) {

	r.POST(path.Join(API, "post", "document", "add-metrics"), func(c *gin.Context) {
		bytes, err := io.ReadAll(c.Request.Body)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "failure",
			})
			return
		}

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
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "failure",
			})
		}

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
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to parse",
			})
		}

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
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to parse",
			})
		}

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

}
