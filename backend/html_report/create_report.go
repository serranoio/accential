package html_report

import (
	"backend/lexer"
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type Document struct {
	HydratedJS  string
	HydratedCSS string
}

func createDocument() *Document {
	return &Document{
		HydratedJS:  "",
		HydratedCSS: "",
	}

}

func getFileNames() []string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fullPath := path.Join(dir, "html_report", "ui", "dist", "assets")

	var a []string
	filepath.WalkDir(fullPath, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		a = append(a, s)
		return nil
	})

	return a

}

func getFile(lookForName string, fileNames []string) string {
	for i := 0; i < len(fileNames); i++ {
		if strings.Contains(fileNames[i], lookForName) {
			return fileNames[i]
		}
	}

	log.Panicln("did not find name")
	return "failed"
}

func (d *Document) hydrateAssets() {
	names := getFileNames()

	jsBytes, err := os.ReadFile(getFile(".js", names))
	if err != nil {
		log.Panic(err)
	}
	cssBytes, err := os.ReadFile(getFile(".css", names))
	if err != nil {
		log.Panic(err)
	}

	d.HydratedCSS = string(cssBytes)
	d.HydratedJS = string(jsBytes)
}

func createReport(statistics *lexer.Statistics) {

	for _, v := range statistics.Mowc.TotalAssets {
		fmt.Println(v)
	}
	for _, v := range statistics.Mowc.TotalLiabilities {
		fmt.Println(v)
	}
	for _, v := range statistics.Mowc.WorkingCapital {
		fmt.Println(v)
	}

	document := createDocument()

	document.hydrateAssets()

	fmt.Println("Done")
	var byteBuf bytes.Buffer // execute template on byte buffer

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fullPath := path.Join(dir, "html_report", "templates", "index.gohtml")

	tmpl := template.Must(template.ParseFiles(fullPath))

	tmpl.Execute(&byteBuf, document)

	os.WriteFile("report.html", byteBuf.Bytes(), 0664)
}
