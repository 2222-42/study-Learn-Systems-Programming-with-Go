package main

import "fmt"

func makeSlices() {
	// 既存の配列を参照
	a := [4]int{0, 1, 2, 3}
	b := a[:]
	fmt.Println(&b[0], len(b), cap(b))
	c := a[1:3]
	fmt.Println(&c[0], len(c), cap(c))

	//何も参照しない
	var d []int
	fmt.Println(len(d), cap(d)) // &d[]はpanicする

	//スライスと裏の配列を同時に作成する
	e := []int{0, 1, 2, 3}
	fmt.Println(&e[0], len(e), cap(e))

	f := make([]int, 4)
	fmt.Println(&f[0], len(f), cap(f))

	g := make([]int, 4, 8)
	fmt.Println(&g[0], len(g), cap(g))
}

func makeSureOfMemory() {
	s := make([]int, 1, 2)
	fmt.Println(&s[0], len(s), cap(s))

	s = append(s, 1)
	fmt.Println(&s[0], len(s), cap(s))

	// cap()の余裕がなくなり、新たな配列が確保され、アドレスも変わる。
	s = append(s, 2)
	fmt.Println(&s[0], len(s), cap(s))
}

func main() {
	makeSlices()
	makeSureOfMemory()
}
