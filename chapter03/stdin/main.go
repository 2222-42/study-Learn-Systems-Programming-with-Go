package main

import (
	"fmt"
	"io"
	"os"
)

// 3.4.1
// nonblockingにするためにgoroutineを使うのは13章
func main() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("sized=%d input='%s'\n", size, string(buffer))
	}
}
