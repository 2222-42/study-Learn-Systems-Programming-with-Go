package main

import (
	"io"
	"os"
)

func main() {
	// OpenもCloseもOpenFile()を使い同じシステムコールを読んでいる
	file, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
