package main

import (
	"io"
	"os"
)

// 2.4.6
func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriter example\n")
}
