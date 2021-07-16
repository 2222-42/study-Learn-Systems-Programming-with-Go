package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// 6.4
// 6.5.2
// 6.6.2
func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil

	for {
		var err error

		if conn == nil {
			conn, err = net.Dial("tcp", ":8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		request, err := http.NewRequest("GET", "http://localhost:8888", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}

		err = request.Write(conn)
		if err != nil {
			panic(err)
		}

		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))
		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}
