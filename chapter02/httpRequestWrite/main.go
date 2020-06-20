package main

import (
	"net/http"
	"os"
)

func main() {
	request, err := http.NewRequest("GET", "http://acii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "hoge")
	request.Write(os.Stdout)
}
