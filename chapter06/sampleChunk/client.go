package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

// 6.8.2
func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	// 6.7.1
	request.Header.Set("Accept-Encoding", "gzip")
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		fmt.Println("Retry")
		conn = nil
	}

	dump, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))

	if len(response.TransferEncoding) < 1 || response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}

	for {
		sizeStr, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		size, err := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}

		line := make([]byte, int(size))
		if _, err := io.ReadFull(reader, line); err != nil {
			panic(err)
		}
		if _, err := reader.Discard(2); err != nil {
			panic(err)
		}
		fmt.Printf("  %d bytes: %s\n", size, string(line))
	}

	conn.Close()
}
