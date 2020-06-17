package main

import "os"

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	// Write() が受け取るのはバイト列なので、変換する。
	file.Write([]byte("os.File example\n"))
	file.Close()
}
