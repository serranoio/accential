package api

import (
	"backend/database"
	"backend/lexer"
	helpers "backend/rabbitmq"
	"encoding/json"
	"net/http"

	// "encoding/base64"
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

var msgs <-chan amqp091.Delivery

func listenForHtmlReport(communication *helpers.Communication) {
	communication.AddConsumingEQ(helpers.EQ_HTML_REPORT, "topic")

	msgs = communication.ConsumeEQ(helpers.EQ_HTML_REPORT)

	// go func() {
	// 	for d := range msgs {
	// 		log.Printf("api: Received a message: %s", d.Body)
	// 		time.Sleep(0 * time.Second)
	// 	}
	// }()

	// log.Printf("api [*] Waiting for messages. To exit press CTRL+C")
}

func InitAPI() {
	database.InitDatabase()

	communication := helpers.CreateCommunication()
	defer communication.Channel.Close()
	defer communication.Connection.Close()

	go listenForHtmlReport(communication)

	time.Sleep(1 * time.Second)

	communication.AddPublshingEQ(helpers.EQ_PDF, "topic")
	defer communication.Context.Cancel()

	temp(communication)

	createApi(communication)

}

var Document []byte

func temp(communication *helpers.Communication) {

	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path.Join(currentWorkingDirectory, "10-k-example.html"))

	helpers.FailOnError(err, "file open")

	bytes, err := io.ReadAll(file)

	helpers.FailOnError(err, "file read")

	Document = bytes

	// send to lexer
	communication.PublishToEQ(helpers.EQ_PDF, bytes)
	// send directly to create_report

}

const API = "api"

func createApi(communication *helpers.Communication) {

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

	r.POST(path.Join(API, "document"), func(c *gin.Context) {
		bytes, err := io.ReadAll(c.Request.Body)
		helpers.FailOnError(err, "failed to parse")

		communication.PublishToEQ(helpers.EQ_PDF, bytes)

		// at the end when we receive everything

		delivery := <-msgs // we wait for channel to finish

		c.JSON(http.StatusOK, gin.H{
			"document is done": string(delivery.Body),
		})
		log.Println("Message sent")
	})

	r.POST(path.Join(API, "metric", "add"), func(c *gin.Context) {
		bytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"failed to parse": "bad",
			})
		}

		metric := &lexer.Metric{}

		err = json.Unmarshal(bytes, &metric)

		database.Db.Create(metric)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"failed to parse": "bad",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.GET(path.Join(API, "metric", "get"), func(c *gin.Context) {

		metrics := []*lexer.Metric{}

		database.Db.Find(&metrics)

		c.JSON(http.StatusOK, gin.H{
			"message": metrics,
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080
	log.Println("listen and serve on 0.0.0.0:8080")
}
