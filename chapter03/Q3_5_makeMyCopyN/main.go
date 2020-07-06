package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func myCopyN(w io.Writer, r io.Reader, n int64) (int64, error) {
	reader := io.LimitReader(r, n)
	written, err := io.Copy(w, reader)
	if written == n {
		return n, nil
	}
	if written < n && err == nil {
		err = io.EOF
	}
	return written, err
}

func main() {
	r := strings.NewReader("123456789012345")
	result1, err1 := myCopyN(os.Stdout, r, 9)
	fmt.Printf("result: %d, err: %v\n", result1, err1)
	result2, err2 := myCopyN(os.Stdout, r, 100)
	fmt.Printf("result: %d, err: %v\n", result2, err2)
}
