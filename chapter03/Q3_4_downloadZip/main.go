package main

import (
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	// file を開く
	file, err := os.Open("sample.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// fileをwのhttp.ResponseWriterに書き込む。
	io.Copy(w, file)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
