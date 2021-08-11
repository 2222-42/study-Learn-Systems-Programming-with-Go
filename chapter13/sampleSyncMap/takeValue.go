package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := &sync.Map{}
	smap.Store("hello", "world")
	smap.Store(1, 2)
	smap.Delete("test")
	smap.LoadOrStore(1, 3)
	smap.LoadOrStore(2, 4)
	value, ok := smap.Load("hello")
	fmt.Printf("key=%v, value=%v, exists?=%v\n", "hello", value, ok)
	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("key=%v, value=%v\n", key, value)
		return true
	})
}
