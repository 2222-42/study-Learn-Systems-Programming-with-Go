package main

import "os"

func main() {
	os.Link("oldFile.txt", "newFile.txt")

	os.Symlink("oldFile.txt", "newFile-symLink.txt")

	link, err := os.Readlink("newFile-symLink.txt")
	if err != nil {
		panic(err)
	}
	println(link)
}
