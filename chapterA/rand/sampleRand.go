package main

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	mrand "math/rand"
	"time"
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

	// set seed
	mrand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		fmt.Println(mrand.Float64())
	}

	// create source from seed
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	src := mrand.NewSource(seed.Int64())
	rng := mrand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(rng.Float64())
	}
}
