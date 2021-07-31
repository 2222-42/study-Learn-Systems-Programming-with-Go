package main

import (
	"fmt"
	"os"
	"time"
)

const fileName = "setting.txt"

func main() {
	os.Chmod(fileName, 0644)
	os.Chown(fileName, os.Getuid(), os.Getegid())
	os.Chtimes(fileName, time.Now(), time.Now())
	info, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("FileInfo")
	fmt.Printf("  ファイル名: %v\n", info.Name())
	fmt.Printf("  サイズ: %v\n", info.Size())
	fmt.Println("Mode()")
	fmt.Printf("  ディレクトリ? %v\n", info.Mode().IsDir())
	fmt.Printf("  読み書き可能な通常ファイル? %v\n", info.Mode().IsRegular())
	fmt.Printf("  Unixのファイルアクセス権限ビット %o\n", info.Mode().Perm())
	fmt.Printf("  モードのテキスト表現 %v\n", info.Mode().String())
}
