package parser

import (
	helpers "backend/rabbitmq"
	"log"
	"time"
)

func listenForPdfFile(communication *helpers.Communication) {
	communication.AddConsumingEQ(helpers.EQ_PDF, "topic")
	// communication.AddPublshingEQ(helpers.EQ)

	msgs := communication.ConsumeEQ(helpers.EQ_PDF)
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("parser: Received a message: %s", d.Body)
			time.Sleep(1 * time.Second)
			communication.PublishToEQ(helpers.EQ_PARSED_PDF, []byte("Hello from parser!"))
		}

	}()

	log.Printf("[*] parser Waiting for messages. To exit press CTRL+C")
	<-forever
}

func InitParser() {
	communication := helpers.CreateCommunication()
	defer communication.Connection.Close()
	defer communication.Channel.Close()

	communication.AddPublshingEQ(helpers.EQ_PARSED_PDF, "topic")
	defer communication.Context.Cancel()

	listenForPdfFile(communication)

}
