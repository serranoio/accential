package main

import (
	lexer "backend/lexer"
	parser "backend/parser"

	converter "backend/converter"
	htmlReport "backend/html_report"

	"time"
)

func main() {

	go htmlReport.BuildReport() // consumer
	go converter.Convert()      // consumer

	time.Sleep(2 * time.Second)

	parser.Parse()   // publisher
	lexer.Tokenize() // publisher
}
