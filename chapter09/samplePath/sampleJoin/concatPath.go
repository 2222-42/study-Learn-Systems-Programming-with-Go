package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func concat() {
	fmt.Printf("Temp File Path : %s\n", filepath.Join(os.TempDir(), "temp.txt"))
}

func split() {
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	fmt.Printf("Dir: %s, Name: %s\n", dir, name)
}

func main() {
	concat()
	split()
}
