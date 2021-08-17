package main

import "fmt"

type Struct struct {
	param string
}

func main() {
	var a int = 10
	var b *Struct = new(Struct)
	var c Struct = Struct{"param"}
	var d *Struct = &Struct{"param"}

	e := [4]int{1, 2, 3, 4}
	f := make([]int, 4, 8)
	g := make(chan string)
	h := make(chan string, 10)
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v", a, b, c, d, e, f, g, h)
}

//  go build -gcflags -m sample.go
