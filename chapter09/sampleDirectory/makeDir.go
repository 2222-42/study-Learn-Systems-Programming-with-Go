package main

import (
	"io"
	"os"
)

func create(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := io.WriteString(file, "New file content\n"); err != nil {
		panic(err)
	}
}

func rename(fileName, newName string) {
	if err := os.Rename(fileName, newName); err != nil {
		panic(err)
	}
}

func removeAll(path string) {
	if err := os.RemoveAll(path); err != nil {
		panic(err)
	}
}

func main() {
	if err := os.Mkdir("setting", 0755); err != nil {
		panic(err)
	}
	if err := os.MkdirAll("setting/myapp/networksettings", 0755); err != nil {
		panic(err)
	}
	if err := os.Remove("setting/myapp/networksettings"); err != nil {
		panic(err)
	}
	removeAll("setting")

	removeAll("olddir")
	removeAll("newdir")
	create("old_name.txt")
	rename("old_name.txt", "new_name.txt")
	if err := os.Mkdir("olddir", 0755); err != nil {
		panic(err)
	}
	create("olddir/file.txt")
	if err := os.Mkdir("newdir", 0755); err != nil {
		panic(err)
	}
	rename("olddir/file.txt", "newdir/file.txt")
	create("olddir/file.txt")
	//rename("olddir/file.txt", "newdir/")
	//rename("olddir/file.txt", "/tmp/sample.rst")

}
