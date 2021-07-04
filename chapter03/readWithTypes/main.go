package main

import (
	"fmt"
	"strings"
)

var source = "123 1.234 1.0e4 \ntest test2"
var source2 = "123, 1.234, 1.0e4, test test3"
var source3 = `123 1.234 1.0e4 test
test2`

// 3.6.2
func main() {
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	// スペースで区切られたデータを読み込む方法
	if _, err := fmt.Fscan(reader, &i, &f, &g, &s); err != nil {
		panic(err)
	}

	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)

	// カンマ+スペースで区切られたデータの読み込み方法
	reader2 := strings.NewReader(source2)
	if _, err := fmt.Fscanf(reader2, "%v, %v, %v, %v", &i, &f, &g, &s); err != nil {
		panic(err)
	}

	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)

	reader3 := strings.NewReader(source3)
	// 改行でスキャンを停止させる
	if _, err := fmt.Fscanln(reader3, &i, &f, &g, &s); err != nil {
		panic(err)
	}

	fmt.Printf("i=%#v f=%#v g=%#v s=%#v\n", i, f, g, s)
}
