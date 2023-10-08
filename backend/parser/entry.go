package parser

// "encoding/base64"
// "fmt"

func InitParser(htmlReport []byte) []*Table {
	return initTablizer(htmlReport)
}
