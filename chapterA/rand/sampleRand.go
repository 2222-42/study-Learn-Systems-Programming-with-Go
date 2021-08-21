package main

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	mrand "math/rand"
)

func main() {
	a := make([]byte, 20)
	if _, err := mrand.Read(a); err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(a))

	b := make([]byte, 20)
	if _, err := crand.Read(b); err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(b))
}
