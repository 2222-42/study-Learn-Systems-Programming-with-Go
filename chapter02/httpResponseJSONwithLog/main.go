package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}
	writer := io.MultiWriter(w, os.Stdout)

	jsonData, err := json.Marshal(source)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := writer.Write(jsonData); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
