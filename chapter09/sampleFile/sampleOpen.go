package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func create() {
	file, err := os.Create("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

func open() {
	file, err := os.Open("textfile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Read file:")
	io.Copy(os.Stdout, file)
}

func append() {
	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Appended content\n")
}

func remove() {
	if err := os.Remove("textfile.txt"); err != nil {
		panic(err)
	}
}

// osを使う方は、そのままファイル名を指定できる
func truncateOs() {
	if err := os.Truncate("textfile.txt", 20); err != nil {
		log.Println(err)
	}
}

// Fileを使う方は、OpenFileを指定しんまいといけない
func truncateFile() {
	file, err := os.OpenFile("textfile.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//  truncate textfile.txt: invalid argument
	if err := file.Truncate(10); err != nil {
		log.Println(err)
	}

}

func main() {
	create()
	open()
	append()
	open()
	truncateOs()
	open()
	truncateFile()
	open()
	remove()
}
