package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("-----HEADER-----\n")
	content := bytes.NewBufferString("Example o fio.MultiReader\n")
	footer := bytes.NewBufferString("-----FOOTER-----\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
}
