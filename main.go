package main

import (
	"io"
	"io/ioutil"
	"os"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling %+v\n", r);
	var htmlFile string
	newFeature := os.Getenv("NEW_FEATURE")

	if newFeature == "true" {
		htmlFile = "/content/newFeature.html"
	} else {
		htmlFile = "/content/index.html"
	}
	bs, err := ioutil.ReadFile(htmlFile)

	if err != nil {
		fmt.Printf("Couldn't read file.html: %v", err);
		os.Exit(1)
	}

	io.WriteString(w, string(bs[:]))
}

func main() {
	http.HandleFunc("/", index)
	port := ":8000"
	fmt.Printf("Starting to service on port %s\n", port);
	http.ListenAndServe(port, nil)
}
