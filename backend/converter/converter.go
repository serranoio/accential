package converter

import (
	helpers "backend/rabbitmq"
	"log"
)

func listenForTokenizedPdfFile(communication *helpers.Communication) {
	communication.AddConsumingEQ(helpers.EQ_TOKENIZED_PDF, "topic")

	msgs := communication.ConsumeEQ(helpers.EQ_TOKENIZED_PDF)
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Converter: Received a message: %s", d.Body)
			communication.PublishToEQ(helpers.EQ_CONVERTED_DATA, []byte("Hello from Conveter!"))
		}

	}()

	log.Printf("Converter [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func InitConverter() *helpers.Communication {
	communication := helpers.CreateCommunication()
	defer communication.Connection.Close()
	defer communication.Channel.Close()

	communication.AddPublshingEQ(helpers.EQ_CONVERTED_DATA, "topic")
	defer communication.Context.Cancel()

	listenForTokenizedPdfFile(communication)

	return communication

}
