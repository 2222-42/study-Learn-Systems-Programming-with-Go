package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var id int64

func generateId(mutex *sync.Mutex) int64 {
	mutex.Lock()
	defer mutex.Unlock()
	id++
	return id
}

func generateIdWithAtomic(mutex *sync.Mutex) int64 {
	return atomic.AddInt64(&id, 1)
}

func main() {
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateIdWithAtomic(&mutex))
		}()
	}
	// 暫定的な対処
	time.Sleep(1 * time.Second)
}
