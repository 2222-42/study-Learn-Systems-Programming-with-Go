package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	newFile, err := os.Create("sample")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	randReader := rand.Reader
	if _, err := io.CopyN(newFile, randReader, 1024); err != nil {
		panic(err)
	}
}
