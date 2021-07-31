package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open("/")
	if err != nil {
		panic(err)
	}

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			fmt.Printf("[Dir] %s\n", info.Name())
		} else {
			fmt.Printf("[File] %s\n", info.Name())
		}
	}
}
