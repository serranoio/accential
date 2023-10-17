package api

import (
	"backend/search"
	"net/http"
	"os"
	"path"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

// miscellaneous routes

func miscellaneousRoutes(r *gin.Engine) {
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
}
