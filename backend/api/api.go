package api

import (
	helpers "backend/rabbitmq"
	"log"
	"time"
)

func Api() {

	var forever chan struct{}

	go func() {

	}()

	<-forever

}

const API = ""

// func pdf(w http.ResponseWriter, rhelpers.EQ *http.Rhelpers.EQuest) {

// 	// fmt.Fprintf(w, "hello\n")
// }

// func main() {

// 	// http.HandleFunc(path.Join(API, "pdf"), pdf)

// 	// http.ListenAndServe(":3000", nil)

// 	consume()
// }

func listenForHtmlReport(communication *helpers.Communication) {
	communication.AddConsumingEQ(helpers.EQ_HTML_REPORT, "topic")

	msgs := communication.ConsumeEQ(helpers.EQ_HTML_REPORT)
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("api: Received a message: %s", d.Body)
		}
	}()

	log.Printf("api [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func InitAPI() {
	communication := helpers.CreateCommunication()

	defer communication.Channel.Close()
	defer communication.Connection.Close()
	go listenForHtmlReport(communication)

	time.Sleep(1 * time.Second)
	communication.AddPublshingEQ(helpers.EQ_PDF, "topic")
	defer communication.Context.Cancel()

	// var forever chan struct{}

	for {
		time.Sleep(1 * time.Second)
		communication.PublishToEQ(helpers.EQ_PDF, []byte("Hello from front-facing api!"))
	}

	// <-forever
}
