package main

// import "github.com/pdfcpu/pdfcpu/pkg/api";

import (
	helpers "backend/rabbitmq"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

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
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func Consume() {
	receive()
}
