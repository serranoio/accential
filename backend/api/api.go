package api

import (
	helpers "backend/rabbitmq"
	"log"
	"net/http"
	"time"

	"path"

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
	// 		time.Sleep(1 * time.Second)
	// 	}
	// }()

	// log.Printf("api [*] Waiting for messages. To exit press CTRL+C")
}

func InitAPI() {

	communication := helpers.CreateCommunication()
	defer communication.Channel.Close()
	defer communication.Connection.Close()

	go listenForHtmlReport(communication)

	time.Sleep(1 * time.Second)

	communication.AddPublshingEQ(helpers.EQ_PDF, "topic")
	defer communication.Context.Cancel()

	createApi(communication)

}

const API = "api"

func createApi(communication *helpers.Communication) {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		// Access-Control-Allow-Origin
		AllowMethods:     []string{"PUT", "GET", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
		AllowFiles:       true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET(path.Join(API, "document"), func(c *gin.Context) {

		// start call

		// c.Request.Body.data

		communication.PublishToEQ(helpers.EQ_PDF, []byte("Hello from front-facing api!"))

		d := <-msgs // we wait for channel to finish

		c.JSON(http.StatusOK, gin.H{
			"document is done": string(d.Body),
		})
		log.Println("Message sent")
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
