package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]", os.Args[0])
		os.Exit(1)
	}

	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
		return
	} else if err != nil {
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

	internalStat := info.Sys().(*syscall.Stat_t)
	fmt.Printf("OS固有情報 %#v\n", internalStat)
}
