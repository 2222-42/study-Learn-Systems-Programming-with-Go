package main

import (
	"fmt"
	"github.com/edsrzf/mmap-go"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var testData = []byte("0123456789ABCDEF")
	var testPath = filepath.Join(os.TempDir(), "testdata")
	if err := ioutil.WriteFile(testPath, testData, 0644); err != nil {
		panic(err)
	}

	f, err := os.OpenFile(testPath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, err := mmap.Map(f, mmap.RDWR, 0) // 読み書きモードで、メモリ上に展開
	if err != nil {
		panic(err)
	}
	defer m.Unmap() // メモリ上に展開された内容を削除して閉じる

	m[9] = 'X'
	if err := m.Flush(); err != nil { // メモリ上のデータをファイルに保存する
		panic(err)
	}

	fileData, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("original: %s\n", testData)
	fmt.Printf("mmap:     %s\n", m)
	fmt.Printf("file:     %s\n", fileData)
}
