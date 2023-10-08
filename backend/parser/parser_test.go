package parser

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"testing"
)

func TestParser(t *testing.T) {

	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path.Join(currentWorkingDirectory, "practice-real.html"))

	bytes, err := io.ReadAll(file)

	tables := InitParser(bytes)

	fmt.Println(tables)
}
