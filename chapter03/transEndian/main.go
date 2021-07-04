package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// 3.5.2
func main() {
	// 32 ビットのビッグエンディアンのデータ（10000）
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32

	// エンディアンの変換
	if err := binary.Read(bytes.NewReader(data), binary.BigEndian, &i); err != nil {
		panic(err)
	}

	//fmt.Printf("data: %d\n")
	fmt.Printf("data: %d\n", i)
}
