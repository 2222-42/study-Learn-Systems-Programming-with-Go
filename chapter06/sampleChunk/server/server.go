package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// 6.8.2
// 青空文庫 : ごんぎつねより
// http://www.aozora.gr.jp/cards/000121/card628.html
var contents = []string{
	" これは、私わたしが小さいときに、村の茂平もへいというおじいさんからきいたお話です。 ",
	" むかしは、私たちの村のちかくの、中山なかやまというところに小さなお城があって、 ",
	" 中山さまというおとのさまが、おられたそうです。 ",
	" その中山から、少しはなれた山の中に、「ごん狐ぎつね」という狐がいました。 ",
	" ごんは、一人ひとりぼっちの小狐で、しだの一ぱいしげった森の中に穴をほって住んでいました。 ",
	" そして、夜でも昼でも、あたりの村へ出てきて、いたずらばかりしました。 ",
}

func processSession(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept %v\n", conn.RemoteAddr())

	for {
		if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
			panic(err)
		}

		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}

		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		if _, err := fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n")); err != nil {
			panic(err)
		}

		for _, content := range contents {
			bytes := []byte(content)
			if _, err := fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content); err != nil {
				panic(err)
			}
		}
		if _, err := fmt.Fprintf(conn, "0\r\n\r\n"); err != nil {
			panic(err)
		}
	}
}

// 6.5
// 6.6.1
func main() {
	ln, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go processSession(conn)

	}
}
