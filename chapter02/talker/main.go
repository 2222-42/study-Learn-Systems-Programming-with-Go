package main

import "fmt"

// define Interface; this has `Talk()` method
type Talker interface {
	Talk()
}

// declare Structure
type Greeter struct {
	name string
}

// the structure have method which were defined in the interface
// 副作用がないのでレシーバーの型はポインタ型ではない
func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	// declare the variable whose has interface type
	var talker Talker
	// substitute the pointer
	// ここで行っていることは、「初期化パラメータを与えてGreeter 型の構造体のインスタンスを作成し、そのポインタを代入」
	// そして、構造体とインターフェースの互換性をチェックしている。
	talker = &Greeter{"wozozo"}
	talker.Talk()
}
