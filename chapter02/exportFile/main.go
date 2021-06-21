package main

import "os"

// 2.4.1
func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	// Write() が受け取るのはバイト列なので、変換する。
	file.Write([]byte("os.File example\n"))
	file.Close()
}
