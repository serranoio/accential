package api

import (
	"fmt"
	"net/http"
	"path"
)

func Api() {

	var forever chan struct{}

	go func() {

	}()

	<-forever

}

const API = ""

func pdf(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func main() {

	http.HandleFunc(path.Join(API, "pdf"), pdf)

	http.ListenAndServe(":3000", nil)
}
