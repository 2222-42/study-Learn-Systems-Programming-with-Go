package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Printf("os open, err: %v", err)
		return
	}
	defer file.Close()
	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("read all, err: %v", err)
		return
	}
	fmt.Printf("read all: %v", string(buffer))

	file2, err := os.Open("main.go")
	if err != nil {
		fmt.Printf("os open, err: %v", err)
		return
	}
	defer file2.Close()
	buffer2 := make([]byte, 4)
	size, err := io.ReadFull(file2, buffer2)
	if err != nil {
		fmt.Printf("read full, err: %v", err)
		return
	}
	fmt.Printf("read full: %v, size: %v", string(buffer2), size)

}
