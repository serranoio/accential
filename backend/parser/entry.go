package parser

func InitParser(htmlReport []byte) []*Table {
	return initTablizer(htmlReport)
}
