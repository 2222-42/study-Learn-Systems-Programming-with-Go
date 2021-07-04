package main

import (
	"io"
	"log"
	"os"
	"strings"
)

// 3.5.1
func main() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	if _, err := io.Copy(os.Stdout, sectionReader); err != nil {
		log.Printf("err: %v", err)
	}
}
