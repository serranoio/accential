package main

import (
	lexer "backend/lexer"
	parser "backend/parser"

	converter "backend/converter"
	html_report "backend/html_report"

	"backend/api"
)

func main() {

	go parser.InitParser()
	go lexer.InitLexer()
	go converter.InitConverter()
	go html_report.InitHtmlReport()

	api.InitAPI()
}
