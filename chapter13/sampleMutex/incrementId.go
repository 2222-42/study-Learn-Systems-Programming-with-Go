package main

import (
	"fmt"
	"sync"
	"time"
)

var id int

func generateId(mutex *sync.Mutex) int {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	return id
}

func main() {
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex))
		}()
	}
	// 暫定的な対処
	time.Sleep(1 * time.Second)
}
