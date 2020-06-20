package main

import (
	"bufio"
	"os"
)

func main() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
	buffer.WriteString("test")

	buffer2 := bufio.NewWriterSize(os.Stdout, 3)
	buffer2.WriteString("bufio.Writer ")
	buffer2.WriteString("example\n")
	buffer2.WriteString("test\n")
}
