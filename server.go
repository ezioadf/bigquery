package main

import (
	"io"
	"fmt"
	"os"
	"net/http"
	"bigquery/bqimport"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Go!")
	return
}

func bqDataImport(w http.ResponseWriter, r *http.Request) {
	bqimport.BQRead()
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/bigquery", bqDataImport)
	fmt.Println("Starting server on port 3002")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}
	http.ListenAndServe(":"+port, nil)
	return
}