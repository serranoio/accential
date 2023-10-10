package html_report

import (
	"backend/comm"
	"backend/database"
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

type HydrateDocument struct {
	HydratedJS   string
	HydratedCSS  string
	HydratedFont string
	Fun          string
	Document     string
	Metrics      []*comm.Metric
	ID           string
}

func createHydrateDocument() *HydrateDocument {
	return &HydrateDocument{
		HydratedJS:  "",
		HydratedCSS: "",
		Metrics:     nil,
		Fun:         "hello",
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

func (d *HydrateDocument) hydrateAssets() {
	names := getFileNames()

	jsBytes, err := os.ReadFile(getFile(".js", names))
	if err != nil {
		log.Panic(err)
	}
	cssBytes, err := os.ReadFile(getFile(".css", names))
	if err != nil {
		log.Panic(err)
	}

	fontBytes, err := os.ReadFile(getFile(".otf", names))
	if err != nil {
		log.Panic(err)
	}

	d.HydratedCSS = string(cssBytes)
	d.HydratedJS = string(jsBytes)
	d.HydratedFont = string(fontBytes)
}

func createReport(metrics []*comm.Metric, htmlReport []byte) uint {
	document := createHydrateDocument()
	document.hydrateAssets()
	document.Metrics = metrics
	document.Document = string(htmlReport)
	// get id from database document, put it in hydration.
	documentToSave := &comm.Document{}
	database.Db.Save(documentToSave)
	document.ID = fmt.Sprintf("%d", documentToSave.ID)

	// execute template on byte buffer
	var byteBuf bytes.Buffer
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fullPath := path.Join(dir, "html_report", "templates", "index.gohtml")
	tmpl := template.Must(template.ParseFiles(fullPath))
	tmpl.Execute(&byteBuf, document)

	reportBytes := byteBuf.Bytes()
	os.WriteFile("report.html", reportBytes, 0664)
	log.Println("File Written")

	return database.SaveDocument(reportBytes, metrics, documentToSave.ID)
}
