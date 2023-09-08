package main

import (
	"context"
	"time"

	helpers "backend/rabbitmq"
)

//   func Init() *amqp.Connection {

// 	//   return conn;
// 	}

func Publish() {

	conn := helpers.InitConnection()
	defer conn.Close()

	ch := helpers.InitChannel(conn)
	defer ch.Close()

	helpers.InitExchange(ch, "fanout")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := "F OUTTA HEAR!"

	helpers.PublishMessage(ch, ctx, []byte(message))
}
