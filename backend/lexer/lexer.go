package lexer

import (
	"context"
	"time"

	helpers "backend/rabbitmq"
)

//   func Init() *amqp.Connection {

// 	//   return conn;
// 	}

func Tokenize() {

	conn := helpers.InitConnection()
	defer conn.Close()

	ch := helpers.InitChannel(conn)
	defer ch.Close()

	helpers.InitExchange(ch, "fanout")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := "Hello from lexer!"

	helpers.PublishMessage(ch, ctx, []byte(message))
}
