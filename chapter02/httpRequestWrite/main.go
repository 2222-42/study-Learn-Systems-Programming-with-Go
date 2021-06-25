package main

import (
	"net/http"
	"os"
)

// 2.4.7
func main() {
	request, err := http.NewRequest("GET", "http://acii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "hoge")
	request.Write(os.Stdout)
}
