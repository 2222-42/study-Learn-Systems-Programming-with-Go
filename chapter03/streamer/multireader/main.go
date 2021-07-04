package main

import (
	"bytes"
	"io"
	"os"
)

// 3.7
func main() {
	header := bytes.NewBufferString("-----HEADER-----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("-----FOOTER-----\n")

	reader := io.MultiReader(header, content, footer)
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		panic(err)
	}
}
