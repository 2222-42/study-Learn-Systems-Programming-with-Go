package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}

	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		execPath := filepath.Join(path, os.Args[1])
		_, err := os.Stat(execPath)
		if !os.IsNotExist(err) {
			fmt.Println(execPath)
			return
		}
	}
	os.Exit(1)
}
