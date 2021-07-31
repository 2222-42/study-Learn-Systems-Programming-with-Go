package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))
	absPath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(absPath)
	relPath, _ := filepath.Rel("/usr/local/go/src", "/usr/local/go/src/path/filepath/path.go")
	fmt.Println(relPath)
}
