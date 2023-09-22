package api

import (
	helpers "backend/rabbitmq"

	// "encoding/base64"
	"io"
	"log"
	"os"
	"path"
	"time"

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

	communication := helpers.CreateCommunication()
	defer communication.Channel.Close()
	defer communication.Connection.Close()

	go listenForHtmlReport(communication)

	time.Sleep(1 * time.Second)

	communication.AddPublshingEQ(helpers.EQ_PDF, "topic")
	defer communication.Context.Cancel()

	temp(communication)

	// createApi(communication)

}

func temp(communication *helpers.Communication) {

	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path.Join(currentWorkingDirectory, "10-k-example.html"))

	helpers.FailOnError(err, "file open")

	bytes, err := io.ReadAll(file)

	helpers.FailOnError(err, "file read")

	forever := make(chan bool)

	communication.PublishToEQ(helpers.EQ_PDF, bytes)

	<-forever
}

const API = "api"

// func createApi(communication *helpers.Communication) {

// 	r := gin.Default()
// 	r.Use(cors.New(cors.Config{
// 		AllowAllOrigins: true,
// 		// Access-Control-Allow-Origin
// 		AllowMethods:     []string{"PUT", "GET", "PATCH"},
// 		AllowHeaders:     []string{"Origin"},
// 		AllowCredentials: true,
// 		AllowFiles:       true,
// 		MaxAge:           12 * time.Hour,
// 	}))

// 	r.POST(path.Join(API, "document"), func(c *gin.Context) {
// 		bytes, err := io.ReadAll(c.Request.Body)
// 		helpers.FailOnError(err, "failed to parse")

// 		communication.PublishToEQ(helpers.EQ_PDF, bytes)

// 		// at the end when we receive everything

// 		delivery := <-msgs // we wait for channel to finish

// 		c.JSON(http.StatusOK, gin.H{
// 			"document is done": string(delivery.Body),
// 		})
// 		log.Println("Message sent")
// 	})

// 	r.Run() // listen and serve on 0.0.0.0:8080
// }
