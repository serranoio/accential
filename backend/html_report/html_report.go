package html_report

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

	q := helpers.DeclareQueue(ch, "name")

	helpers.BindQueue(q, ch, "logs")

	msgs := helpers.Consume(q, ch)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("html_report: Received a message: %s", d.Body)
		}
	}()

	log.Printf("html_report [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func BuildReport() {

	receive()
}
