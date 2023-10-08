package parser

import (
	"backend/comm"
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

func (t *Table) addTr(name string) {
	t.Tr[name] = []string{}
}

func (t *Table) addTableName(name []byte) {
	t.TableName = string(name)
}

func (t *Table) printTable() {
	for k, _ := range t.Tr {
		fmt.Println("key:", k)
	}
}

func (t *Table) addTd(currentTr string, content string) {
	t.Tr[currentTr] = append(t.Tr[currentTr], content)
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

	tnUnclean, _ := tokenizer.TagName()
	tn := comm.CleanString(tnUnclean)
	if tt == html.EndTagToken && includes(findTagNames, tn) {
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

	tnUnclean, _ := tokenizer.TagName()
	tn := comm.CleanString(tnUnclean)
	if includes(endTagNames, tn) &&
		tt == html.EndTagToken { // base case

		return nil, errors.New(fmt.Sprintf("end of %s", endTagNames))
	} else if findText && includes(findTagNames, tn) {

		return []byte(""), nil
	} else if includes(findTagNames, tn) && !findText { // found text box

		return getToken(tokenizer, []string{"td"}, true, endTagNames)
	} else if findText { // get contents in textbox
		// find everything within td

		return getContents(tokenizer, []string{"td"}, []byte{}), nil
	}

	return getToken(tokenizer, findTagNames, findText, endTagNames)
}

func recurse(tokenizer *html.Tokenizer, tables []*Table, depth *int) []*Table {
	// it currently goes to the SPAN first because a call to table will take first column
	for { // for every table
		tt := tokenizer.Next()
		tnUnclean, _ := tokenizer.TagName()

		tn := comm.CleanString(tnUnclean)

		if tn == "body" && tt == html.EndTagToken || tt == html.ErrorToken {
			break
		}
		if tn != "table" {
			continue
		}

		var table *Table
		tables, table = newTable(tables)

		for { // for every row
			Tr, err := getToken(tokenizer, []string{"td"}, false, []string{"table"})
			if err != nil {
				break
			}
			trString := comm.CleanString(Tr)
			table.addTr(trString)

			if trString == "totalliabilitiesandequity" {
				fmt.Println("gey")
			}

			for { // for every td
				p, err := getToken(tokenizer, []string{"td"}, false, []string{"tr"})

				if err != nil {
					break
				}

				pString := comm.CleanString(p)
				table.addTd(trString, pString)
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
