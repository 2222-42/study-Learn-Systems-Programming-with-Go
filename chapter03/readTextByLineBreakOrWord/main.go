package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `Line 1
Line 2
Line 3`

// 3.6.1
func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		if err == io.EOF {
			break
		}
	}

	// 終端を気にせず短く各方法
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}

	// デフォルトの改行区切りではなく、単語区切りにしてみるケース
	scanner2 := bufio.NewScanner(strings.NewReader(source))
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		fmt.Printf("%#v\n", scanner2.Text())
	}
}
