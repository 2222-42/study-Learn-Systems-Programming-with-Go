package main

import (
	"fmt"
	"io"
	"os"
)

// 3.2.2
func main() {
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Printf("os open, err: %v\n", err)
		return
	}
	defer file.Close()
	writeSize, err := io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Printf("Copy, err: %v\n", err)
		return
	}
	fmt.Printf("Copy size: %v\n", writeSize)

	file2, err := os.Open("main.go")
	if err != nil {
		fmt.Printf("os open, err: %v\n", err)
		return
	}
	defer file2.Close()

	writeSize2, err := io.CopyN(os.Stdout, file2, 4)
	if err != nil {
		fmt.Printf("CopyN, err: %v\n", err)
		return
	}
	fmt.Printf("CopyN size: %v\n", writeSize2)

	buffer := make([]byte, 1024)
	writeSize3, err := io.CopyBuffer(os.Stdout, file2, buffer)
	if err != nil {
		fmt.Printf("CopyBuffer, err: %v\n", err)
		return
	}
	fmt.Printf("CopyBuffer size: %v\n", writeSize3)

}
