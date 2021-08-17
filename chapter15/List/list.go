package main

import "fmt"

//type vector3D [3]float64
//type color[4]uint8

func main() {
	var a [4]int
	b := [4]int{}
	c := [4]int{0, 1, 2, 3}
	d := [...]int{0, 1, 2, 3}
	fmt.Println(a[1])
	fmt.Println(b[1])
	fmt.Println(c[1])
	fmt.Println(d[1])

	// コンパイル時のエラーとなる。
	// ./list.go:14:15: invalid array index 10 (out of bounds for 4-element array)
	//fmt.Println(a[10])
	//fmt.Println(a[-1])

	// 変数を使う場合、go buildではエラーが出ず、go runでの実行時エラーになる
	// panic: runtime error: index out of range [10] with length 4
	// i := 10
	//fmt.Println(d[i])
}
