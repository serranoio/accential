package rabbitmq

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// both

func InitConnection() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func InitChannel(conn *amqp.Connection) *amqp091.Channel {
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	return ch
}

func InitExchange(ch *amqp091.Channel, exchangeType, exchangeName string) {
	err := ch.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare en exchange")
}

// publisher

// func PublishMessage(ch *amqp091.Channel, ctx context.Context, message []byte) {

// 	err := ch.PublishWithContext(ctx,
// 		"logs", // exchange
// 		"",     // routing key
// 		false,  // mandatory
// 		false,  // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(message),
// 		})

// 	FailOnError(err, "Failed to publish a message")

// 	log.Printf(" [x] Sent %s", message)
// }

func PublishMessage(ch *amqp091.Channel, ctx context.Context, exchangeName string, message []byte) {

	err := ch.PublishWithContext(ctx,
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	FailOnError(err, "Failed to publish a message")
}

// consumer

func DeclareQueue(ch *amqp091.Channel, name string) amqp091.Queue {

	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	return q
}

// Binds queue with exchange
func BindQueue(q amqp091.Queue, ch *amqp091.Channel, exchangeName string) {
	err := ch.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	FailOnError(err, "Failed to bind queue")
}

func Consume(q amqp091.Queue, ch *amqp091.Channel) <-chan amqp091.Delivery {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	return msgs
}

// func ConsumeOnQueue() {

// }

// func PublishOnQueue() {

// }
