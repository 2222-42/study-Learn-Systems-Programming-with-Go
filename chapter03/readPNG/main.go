package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// 3.5.3
func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	if _, err := file.Seek(8, 0); err != nil {
		panic(err)
	}

	var offset int64 = 8

	for {
		var length int32

		// ここで現在位置は長さを読み終わった箇所になる
		if err := binary.Read(file, binary.BigEndian, &length); err == io.EOF {
			break
		}

		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))

		offset, _ = file.Seek(int64(length+8), 1)
	}

	return chunks
}

func dumpChunk(chunk io.Reader) {
	var length int32
	if err := binary.Read(chunk, binary.BigEndian, &length); err != nil {
		panic(err)
	}

	buffer := make([]byte, 4)
	if _, err := chunk.Read(buffer); err != nil {
		panic(err)
	}

	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func main() {
	file, err := os.Open("Lenna_(test_image).png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
