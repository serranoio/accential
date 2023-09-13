package rabbitmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Context struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

type Queue struct {
	QueueName string
	Queue     amqp091.Queue
}

type Communication struct {
	ChannelName  *amqp091.Channel
	ExchangeName string
	Publishers   []string
	Consumers    []string
	Context      Context
	Queue        *Queue
}

// creates Communication... this will create
// a topic exchange
// a queue
// a channel connection

// exchange
func CreateCommunication(conn *amqp.Connection, exchangeName string, queueName string, publishers, consumers []string) *Communication {

	channel := InitChannel(conn)
	InitExchange(channel, "topic", exchangeName)
	q := DeclareQueue(channel, queueName)
	BindQueue(q, channel, exchangeName)

	// CREATE CONTEXT HERE TOO
	return &Communication{
		ChannelName:  channel,
		ExchangeName: exchangeName,
		Publishers:   publishers,
		Consumers:    consumers,
		Queue: &Queue{
			QueueName: queueName,
			Queue:     q,
		},
	}
}

//

func (c *Communication) AddPublisherName(name string) {
	c.Publishers = append(c.Publishers, name)
}

func (c *Communication) PublishMessage(message []byte) {
	PublishMessage(c.ChannelName, c.Context.Ctx, message)
}

func (c *Communication) AddConsumerName(name string) {
	c.Consumers = append(c.Consumers, name)
}

func (c *Communication) ConsumeMessage(name string) {
	Consume(c.Queue.Queue, c.ChannelName)
}

// func listen
