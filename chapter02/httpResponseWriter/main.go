package main

import (
	"io"
	"net/http"
)

//2.4.5
func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponseWriter sample")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
