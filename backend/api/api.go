package api

import (
	finance "backend/financizer"
	"backend/html_report"
	"backend/parser"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type FileAndMetric struct {
	File   string
	Metric string
}

func InitAPI() {
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
		// AllowAllOrigins: true,
		AllowOrigins:     []string{"https://accential.pages.dev"},
		AllowMethods:     []string{"PUT", "GET", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
		AllowFiles:       true,
		MaxAge:           12 * time.Hour,
	}))

	metricRoutes(r)
	documentRoutes(r)
	miscellaneousRoutes(r)

	r.Run(":8080")
}
