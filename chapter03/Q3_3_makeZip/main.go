package main

import (
	"archive/zip"
	"os"
	"strings"
)

func main() {
	//archive/zip パッケージを使ってzip ファイルを作成
	//、zip ファイルの書き込み用の構造体
	file, err := os.Create("sample.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	//上記の例では、newfile.txt という実際のファイルが、最初に作った出力先の ファイルfileへと圧縮されます
	// 圧縮していない
	if _, err := zipWriter.Create("dir/test2.txt"); err != nil {
		panic(err)
	}

	writer, err := zipWriter.Create("dir/test1.txt")
	if err != nil {
		panic(err)
	}

	//文字列strings.Reader を使ってzip ファイルに書き込む必要がある。
	str := strings.NewReader("sample_message for strings.NewReader")
	if _, err := str.WriteTo(writer); err != nil {
		panic(err)
	}
}
