package converter

import (
	"log"

	helpers "backend/rabbitmq"
)

//   func Init() *amqp.Connection {

// 	//   return conn;
// 	}

func receive() {
	conn := helpers.InitConnection()
	defer conn.Close()

	ch := helpers.InitChannel(conn)
	defer ch.Close()

	q := helpers.DeclareQueue(ch)

	helpers.BindQueue(q, ch, "logs")

	msgs := helpers.Consume(q, ch)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Converter: Received a message: %s", d.Body)
		}
	}()

	log.Printf("Converter [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func Convert() {

	receive()
}
