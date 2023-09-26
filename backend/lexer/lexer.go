package lexer

import (
	"backend/parser"
	helpers "backend/rabbitmq"
	"bytes"
	"encoding/gob"
	"log"
)

func receiveTable(stream []byte) []*parser.Table {

	enc := gob.NewDecoder(bytes.NewReader(stream))

	var tables []*parser.Table

	err := enc.Decode(&tables)

	if err != nil {
		log.Panicln("did not decode")
	}

	return tables
}

func sendStatistics(s *Statistics) []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(&s)

	if err != nil {
		log.Panicln("did not encode")
	}

	return network.Bytes()
}

func listenForParsedPdfFile(communication *helpers.Communication) {
	communication.AddConsumingEQ(helpers.EQ_PARSED_PDF, "topic")

	msgs := communication.ConsumeEQ(helpers.EQ_PARSED_PDF)
	var forever chan struct{}

	go func() {
		for d := range msgs {
			tables := receiveTable(d.Body)

			statistics := processTables(tables)

			communication.PublishToEQ(helpers.EQ_CONVERTED_DATA, sendStatistics(statistics))
		}

	}()

	log.Printf("lexer [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func InitLexer() *helpers.Communication {
	communication := helpers.CreateCommunication()
	defer communication.Connection.Close()
	defer communication.Channel.Close()

	communication.AddPublshingEQ(helpers.EQ_CONVERTED_DATA, "topic")
	defer communication.Context.Cancel()

	listenForParsedPdfFile(communication)

	return communication
}
