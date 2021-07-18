package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

// 6.9.2
func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	var conn net.Conn = nil
	var err error
	requests := make([]*http.Request, 0, len(sendMessages))

	conn, err = net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i := 0; i < len(sendMessages); i++ {
		lastMessage := i == len(sendMessages)-1
		request, err := http.NewRequest("GET", "http://localhost:8888?="+sendMessages[i], nil)
		if err != nil {
			panic(err)
		}
		if lastMessage {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("Connection", "keep-alive")
		}

		if err = request.Write(conn); err != nil {
			panic(err)
		}

		fmt.Println("send: ", sendMessages[i])
		requests = append(requests, request)
	}

	reader := bufio.NewReader(conn)
	for _, request := range requests {
		response, err := http.ReadResponse(reader, request)
		fmt.Println(request.URL.Query())
		if err != nil {
			panic(err)
		}

		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	}
}
