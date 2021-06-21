package main

import (
	"bytes"
	"fmt"
	"io"
)

// 2.4.3
// bytes.Bufferには読み込みの抽象化のio.Readerの昨日もある。
func main() {
	var buffer bytes.Buffer
	// ためておいて
	buffer.Write([]byte("bytes.Buffer example\n"))
	// WriteStringはio.Writerのメソッドではない
	buffer.WriteString("bytes.Buffer example\n")
	io.WriteString(&buffer, "bytes.Buffer example\n")
	// まとめて結果を受け取る
	fmt.Println(buffer.String())
}
