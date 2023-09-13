package rabbitmq

import (
	"context"
	"time"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

const EQ_PDF = "eq-pdf"
const EQ_NOTIFICATION = "eq-notification"
const EQ_PARSED_PDF = "eq-parsed-pdf"
const EQ_TOKENIZED_PDF = "eq-tokenized-pdf"
const EQ_CONVERTED_DATA = "eq-converted-data"
const EQ_HTML_REPORT = "eq-html-report"

type Context struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

// this has the name of the exchange and name of the queue
type EQPairing struct {
	Queue amqp091.Queue
	Name  string
}

type Communication struct {
	Connection   *amqp.Connection
	Channel      *amqp091.Channel
	Context      *Context
	PublishingEQ map[string]EQPairing
	ConsumingEQ  map[string]EQPairing
}

// creates Communication... this will create
// a topic exchange
// a queue
// a channel connection

// exchange
func CreateCommunication() *Communication {
	conn := InitConnection()

	channel := InitChannel(conn)
	return &Communication{
		Connection:   conn,
		Channel:      channel,
		Context:      nil,
		PublishingEQ: make(map[string]EQPairing),
		ConsumingEQ:  make(map[string]EQPairing),
	}
}

func fuzeEQ(c *Communication, name, exchangeType string) amqp.Queue {
	q := DeclareQueue(c.Channel, name)
	InitExchange(c.Channel, exchangeType, name)
	BindQueue(q, c.Channel, name)

	return q
}

// creates context. make sure to defer canceling
func (c *Communication) AddPublshingEQ(name, exchangeType string) {
	q := fuzeEQ(c, name, exchangeType)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	c.Context = &Context{
		Ctx:    ctx,
		Cancel: cancel,
	}

	c.PublishingEQ[name] = EQPairing{
		Queue: q,
		Name:  name,
	}

	// log.Println("Added publishing queue with name", name)
}

func (c *Communication) AddConsumingEQ(name, exchangeType string) {
	q := fuzeEQ(c, name, exchangeType)

	c.ConsumingEQ[name] = EQPairing{
		Queue: q,
		Name:  name,
	}
}

func (c *Communication) PublishToEQ(name string, message []byte) {
	PublishMessage(c.Channel, c.Context.Ctx, name, message)
}

func (c *Communication) ConsumeEQ(name string) <-chan amqp091.Delivery {
	return Consume(c.ConsumingEQ[name].Queue, c.Channel)
}

//

// func (c *Communication) AddPublisherName(name string) {
// 	c.Publishers = append(c.Publishers, name)
// }

// func (c *Communication) PublishMessage(message []byte) {
// 	PublishMessage(c.Channel, c.Context.Ctx, message)
// }

// // func (c *Communication) AddConsumerName(name string) {
// // 	c.Consumers = append(c.Consumers, name)
// // }

// func (c *Communication) ConsumeMessage() <-chan amqp091.Delivery {
// 	return Consume(c.Queue.Queue, c.Channel)
// }

// func listen
