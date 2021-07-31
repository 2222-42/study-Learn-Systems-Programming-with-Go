package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func expandEnv() {
	path := os.ExpandEnv("${GOPATH}/src/github.com")
	fmt.Println(path)
}

func printHomeDirectory() {
	fmt.Println(os.UserHomeDir())
}

func clean2(path string) string {
	if len(path) > 1 && path[0:2] == "~/" {
		my, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = my.HomeDir + path[1:]
	}
	path = os.ExpandEnv(path)
	return filepath.Clean(path)
}

func main() {
	expandEnv()
	printHomeDirectory()
	fmt.Println(clean2("~/Documents/memo/"))
}
