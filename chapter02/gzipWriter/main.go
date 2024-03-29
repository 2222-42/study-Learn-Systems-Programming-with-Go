package main

import (
	"compress/gzip"
	"io"
	"os"
)

// 2.4.6
func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	defer writer.Close()
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "gzip.Writer example\n")
}
