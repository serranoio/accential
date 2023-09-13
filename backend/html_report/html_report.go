package html_report

import (
	helpers "backend/rabbitmq"
	"log"
)

func listenForConvertedData(communication *helpers.Communication) {
	communication.AddConsumingEQ(helpers.EQ_CONVERTED_DATA, "topic")

	msgs := communication.ConsumeEQ(helpers.EQ_CONVERTED_DATA)
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("html-report: Received a message: %s", d.Body)
			communication.PublishToEQ(helpers.EQ_HTML_REPORT, []byte("Hello from html-report!"))
		}

	}()

	log.Printf("html-report [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func PublishHtmlReport(communication *helpers.Communication) {

}

func InitHtmlReport() *helpers.Communication {
	communication := helpers.CreateCommunication()
	defer communication.Connection.Close()
	defer communication.Channel.Close()

	communication.AddPublshingEQ(helpers.EQ_HTML_REPORT, "topic")
	defer communication.Context.Cancel()

	listenForConvertedData(communication)

	return communication
}
