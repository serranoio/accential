package parser

import (
	helpers "backend/rabbitmq"
	"bytes"
	"fmt"

	// "encoding/base64"
	// "fmt"
	"log"

	"golang.org/x/net/html"
)

func listenForPdfFile(communication *helpers.Communication) {
	communication.AddConsumingEQ(helpers.EQ_PDF, "topic")
	// communication.AddPublshingEQ(helpers.EQ)

	msgs := communication.ConsumeEQ(helpers.EQ_PDF)
	var forever chan struct{}

	go func() {
		d := <-msgs

		ioReader := bytes.NewReader(d.Body)

		tokenizer := html.NewTokenizer(ioReader)

		depth := 0

	out:
		for {
			tt := tokenizer.Next()
			// break
			switch tt {
			case html.ErrorToken:
				tokenizer.Err()
				break out
			case html.TextToken:
				if depth > 0 {
					// emitBytes should copy the []byte it receives,
					// if it doesn't process it immediately.

					fmt.Printf("%s", tokenizer.Text())
				}
			case html.StartTagToken, html.EndTagToken:
				tn, _ := tokenizer.TagName()
				if len(tn) == 1 && tn[0] == 'p' {
					if tt == html.StartTagToken {
						depth++
					} else {
						depth--
					}
				}
			}
		}

		communication.PublishToEQ(helpers.EQ_PARSED_PDF, []byte("Hello from parser!"))
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
