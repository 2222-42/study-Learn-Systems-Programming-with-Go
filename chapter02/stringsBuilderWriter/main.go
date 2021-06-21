package main

import (
	"fmt"
	"strings"
)

// 2.4.4
// 書き出し専用
func main() {
	var builder strings.Builder
	builder.Write([]byte("strings.Builder example\n"))
	builder.Write([]byte("strings.Builder example\n"))
	// 読み出しがString()のみ
	fmt.Println(builder.String())
}
