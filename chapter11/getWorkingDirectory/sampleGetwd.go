package main

import "os"

func main() {
	wd, _ := os.Getwd()
	println(wd)
}
