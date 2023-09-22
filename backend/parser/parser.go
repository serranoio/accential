package parser

import (
	helpers "backend/rabbitmq"
	"bytes"
	"encoding/gob"

	// "encoding/base64"
	// "fmt"
	"log"
)

func sendTable(tables []*Table) []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(tables)

	if err != nil {
		log.Panicln("did not encode")
	}

	return network.Bytes()
}

func listenForPdfFile(communication *helpers.Communication) {
	communication.AddConsumingEQ(helpers.EQ_PDF, "topic")
	// communication.AddPublshingEQ(helpers.EQ)

	msgs := communication.ConsumeEQ(helpers.EQ_PDF)
	var forever chan struct{}

	go func() {
		d := <-msgs

		tables := initTablizer(d.Body)
		// parseText(d.Body)

		communication.PublishToEQ(helpers.EQ_PARSED_PDF, sendTable(tables))
		// }

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
