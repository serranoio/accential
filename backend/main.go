package main

import (
	// lexer "backend/lexer"
	// parser "backend/parser"

	// converter "backend/converter"

	"backend/api"
	"backend/converter"
	"backend/html_report"
	"backend/lexer"
	"backend/parser"
)

func main() {

	// consumers
	go parser.InitParser()
	go lexer.InitLexer()
	go converter.InitConverter()
	go html_report.InitHtmlReport()

	api.InitAPI()
	// publishers
	// for {
	// time.Sleep(5 * time.Second)
	// fmt.Println("called")
	// api.PublishPdfFile(comm)

	// }
}
