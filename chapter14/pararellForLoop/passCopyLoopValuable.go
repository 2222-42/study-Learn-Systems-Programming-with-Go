package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []string{
		"cmake ...",
		"cmake . --build Release",
		"cpack",
	}
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		// ループ変数の実体は1つなので、引数として渡しておく必要がある。
		go func(task string) {
			fmt.Println(task)
			wg.Done()
		}(task)
	}
	wg.Wait()
}
