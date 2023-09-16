package lexer

import (
	helpers "backend/rabbitmq"
	"log"
	"time"
)

func listenForParsedPdfFile(communication *helpers.Communication) {
	communication.AddConsumingEQ(helpers.EQ_PARSED_PDF, "topic")

	msgs := communication.ConsumeEQ(helpers.EQ_PARSED_PDF)
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("lexer: Received a message: %s", d.Body)
			time.Sleep(1 * time.Second)
			communication.PublishToEQ(helpers.EQ_TOKENIZED_PDF, []byte("Hello from Tokenizer!"))
		}

	}()

	log.Printf("lexer [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func InitLexer() *helpers.Communication {
	communication := helpers.CreateCommunication()
	defer communication.Connection.Close()
	defer communication.Channel.Close()

	communication.AddPublshingEQ(helpers.EQ_TOKENIZED_PDF, "topic")
	defer communication.Context.Cancel()

	listenForParsedPdfFile(communication)

	return communication
}
