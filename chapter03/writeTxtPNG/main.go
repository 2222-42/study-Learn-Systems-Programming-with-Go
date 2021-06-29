package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

// 3.5.4
func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	if _, err := file.Seek(8, 0); err != nil {
		panic(err)
	}

	var offset int64 = 8

	for {
		var length int32
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
	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		if _, err := chunk.Read(rawText); err != nil {
			panic(err)
		}

		fmt.Println(string(rawText))
	}
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)
	var buffer bytes.Buffer

	// 長さ
	if err := binary.Write(&buffer, binary.BigEndian, int32(len(byteData))); err != nil {
		panic(err)
	}

	// 種類(チャンク名)
	if _, err := buffer.WriteString("tEXt"); err != nil {
		panic(err)
	}

	// データ
	if _, err := buffer.Write(byteData); err != nil {
		panic(err)
	}

	crc := crc32.NewIEEE()
	if _, err := io.WriteString(crc, "tEXt"); err != nil {
		panic(err)
	}

	if err := binary.Write(&buffer, binary.BigEndian, crc.Sum32()); err != nil {
		panic(err)
	}

	return &buffer
}

func readAgain() {
	file, err := os.Open("Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}

func main() {
	file, err := os.Open("Lenna_(test_image).png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFile, err := os.Create("Lenna2.png")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	chunks := readChunks(file)

	if _, err := io.WriteString(newFile, "\x89PNG\r\n\x1a\n"); err != nil {
		panic(err)
	}

	if _, err := io.Copy(newFile, chunks[0]); err != nil {
		panic(err)
	}

	if _, err := io.Copy(newFile, textChunk("ASCII PROGRAMMING++")); err != nil {
		panic(err)
	}

	for _, chunk := range chunks[1:] {
		if _, err := io.Copy(newFile, chunk); err != nil {
			panic(err)
		}
	}
	readAgain()
}
