package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

// 3.7
func main() {
	var buffer bytes.Buffer

	reader := bytes.NewBufferString("Example o fio.TeeReader\n")
	//Reader が読み込んだ内容をバッファにも入れている
	teeReader := io.TeeReader(reader, &buffer)

	// データを読み捨てる
	_, _ = ioutil.ReadAll(teeReader)

	// けどバッファに残ってる
	fmt.Println(buffer.String())
}
