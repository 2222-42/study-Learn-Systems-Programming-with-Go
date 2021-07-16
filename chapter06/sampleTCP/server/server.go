package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

// 6.5
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

		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}

			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))

			response := http.Response{
				Status:           "",
				StatusCode:       200,
				Proto:            "",
				ProtoMajor:       1,
				ProtoMinor:       0,
				Header:           nil,
				Body:             ioutil.NopCloser(strings.NewReader("Hello World\n")),
				ContentLength:    0,
				TransferEncoding: nil,
				Close:            false,
				Uncompressed:     false,
				Trailer:          nil,
				Request:          nil,
				TLS:              nil,
			}
			if err := response.Write(conn); err != nil {
				panic(err)
			}
			conn.Close()
		}()

	}
}
