package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// 6.9.1
// 開始した順に書き出し
// チャンネルに入れた(開始した)順に処理される
func writeToConn(sessionResponses chan chan *http.Response, conn net.Conn) {
	defer conn.Close()
	// 順番に取り出す
	for sessionResponse := range sessionResponses {
		// 選択された仕事が終わるまで待つ。
		response := <-sessionResponse
		if err := response.Write(conn); err != nil {
			panic(err)
		}
		close(sessionResponse)
	}
}

func handleRequest(request *http.Request, sessionResponse chan *http.Response) {
	dump, err := httputil.DumpRequest(request, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
	content := "content is " + string(dump)
	response := &http.Response{
		StatusCode:    200,
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(strings.NewReader(content)),
		ContentLength: int64(len(content)),
	}
	sessionResponse <- response
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	sessionResponses := make(chan chan *http.Response, 50)
	defer close(sessionResponses)
	go writeToConn(sessionResponses, conn)
	reader := bufio.NewReader(conn)

	for {
		if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
			panic(err)
		}

		request, err := http.ReadRequest(reader)
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

		sessionResponse := make(chan *http.Response)
		sessionResponses <- sessionResponse
		go handleRequest(request, sessionResponse)
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
