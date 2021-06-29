package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"time"
)

// 3.7
func main() {
	pr, pw := io.Pipe()
	mw := bufio.NewWriter(pw)
	// TODO: ここはもうちょい非同期なのを自作したい(テキストには書いてなかったので)
	go func() {
		defer pw.Close()

		if _, err := mw.WriteString("-----HEADER-----\n"); err != nil {
			panic(err)
		}

		if err := mw.Flush(); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
		if _, err := mw.WriteString("Example of io.Pipe\n"); err != nil {
			panic(err)
		}

		if err := mw.Flush(); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
		if _, err := mw.WriteString("-----FOOTER-----\n"); err != nil {
			panic(err)
		}

		if err := mw.Flush(); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)

	}()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(pr); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())

}
