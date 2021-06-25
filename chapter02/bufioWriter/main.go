package main

import (
	"bufio"
	"os"
)

// 2.4.6
func main() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
	buffer.WriteString("test\n") //出力されず消える。
	//buffer.Flush()

	buffer2 := bufio.NewWriterSize(os.Stdout, 3) // バッファサイズを指定して自動で呼び出すようにする。
	buffer2.WriteString("bufio.Writer ")
	buffer2.WriteString("example\n")
	buffer2.WriteString("test\n")
	buffer2.WriteString("test\n")
}
