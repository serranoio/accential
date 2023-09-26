package parser

import (
	"bytes"
	"errors"
	"fmt"

	"golang.org/x/net/html"
)

// type RowInfo struct {
// 	RowName string;
// }

type Td []string

type Table struct {
	Tr        map[string]Td // we search for Tr, once we find it, we populate Td.
	TableName string
}

func (t *Table) addTr(name []byte) {
	t.Tr[string(name)] = []string{}
}

func (t *Table) addTableName(name []byte) {
	t.TableName = string(name)
}

func (t *Table) printTable() {
	for k, _ := range t.Tr {
		fmt.Println("key:", k)
	}
}

func (t *Table) addTd(currentTr []byte, content []byte) {
	t.Tr[string(currentTr)] = append(t.Tr[string(currentTr)], string(content))
}

// if we find the word assets in a row, we will then search for corresponding
// numbers in the next columns
// if there exists no assets, DIP.

// adds table to the list of tables
// returns new table
func newTable(tables []*Table) ([]*Table, *Table) {
	t := &Table{
		Tr:        make(map[string]Td),
		TableName: "",
	}

	tables = append(tables, t)

	return tables, t
}

func includes(findTagNames []string, tn string) bool {
	for _, v := range findTagNames {
		if v == tn {
			return true
		}
	}
	return false
}

// leave findTagNames as just p
func getContents(tokenizer *html.Tokenizer, findTagNames []string, allText []byte) []byte {
	tt := tokenizer.Next()

	tn, _ := tokenizer.TagName()
	if tt == html.EndTagToken && includes(findTagNames, string(tn)) {
		return allText
	}

	if tt == html.TextToken {
		allText = append(allText, tokenizer.Text()...)
		return getContents(tokenizer, findTagNames, allText)
	}

	return getContents(tokenizer, findTagNames, allText)
}

func getToken(tokenizer *html.Tokenizer, findTagNames []string, findText bool, endTagNames []string) ([]byte, error) {
	tt := tokenizer.Next()

	tn, _ := tokenizer.TagName()
	if includes(endTagNames, string(tn)) && tt == html.EndTagToken { // base case

		return nil, errors.New(fmt.Sprintf("end of %s", endTagNames))
	} else if includes(findTagNames, string(tn)) && !findText {

		return getToken(tokenizer, []string{"p", "span"}, true, endTagNames)
	} else if findText && includes(findTagNames, string(tn)) {
		// find everything within p

		return getContents(tokenizer, []string{"p"}, []byte{}), nil
	}

	return getToken(tokenizer, findTagNames, findText, endTagNames)
}

func recurse(tokenizer *html.Tokenizer, tables []*Table, depth *int) []*Table {
	// it currently goes to the SPAN first because a call to table will take first column
	for { // for every table
		Tr, err := getToken(tokenizer, []string{"table"}, false, []string{"body"})

		if err != nil {
			break
		}

		var table *Table
		tables, table = newTable(tables)

		table.addTableName(Tr)
		table.addTr(Tr)

		for { // for every td
			p, err := getToken(tokenizer, []string{"td"}, false, []string{"tr"})

			table.addTd(Tr, p)
			if err != nil {
				break
			}
		}

		for { // for every row
			Tr, err = getToken(tokenizer, []string{"tr"}, false, []string{"table"})
			table.addTr(Tr)

			if err != nil {
				break
			}

			for { // for every td
				p, err := getToken(tokenizer, []string{"td"}, false, []string{"tr"})

				table.addTd(Tr, p)
				if err != nil {
					break
				}
			}
		}
	}

	return tables

}

func createTokenizer(htmlDocument []byte) *html.Tokenizer {
	ioReader := bytes.NewReader(htmlDocument)
	tokenizer := html.NewTokenizer(ioReader)

	return tokenizer
}

func initTablizer(htmlDocument []byte) []*Table {
	tokenizer := createTokenizer(htmlDocument)
	tables := []*Table{}
	depth := 0

	tables = recurse(tokenizer, tables, &depth)

	return tables
	// parseText(tokenizer, tables, &depth)

}
