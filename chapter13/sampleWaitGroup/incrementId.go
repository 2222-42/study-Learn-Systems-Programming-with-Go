package main

import (
	"fmt"
	"sync"
)

var id int

func generateId(mutex *sync.Mutex, wg *sync.WaitGroup) int {
	mutex.Lock()
	defer mutex.Unlock()
	defer wg.Done()
	id++
	return id
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("id: %d\n", generateId(&mutex, &wg))
		}()
	}

	wg.Wait()
	fmt.Println("done")
}
