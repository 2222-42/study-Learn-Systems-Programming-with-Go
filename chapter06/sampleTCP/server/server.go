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

		go func() {
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
				content := "Hello World\n"

				response := http.Response{
					Status:           "",
					StatusCode:       200,
					Proto:            "",
					ProtoMajor:       1,
					ProtoMinor:       1,
					Header:           nil,
					Body:             ioutil.NopCloser(strings.NewReader(content)),
					ContentLength:    int64(len(content)),
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
			}
		}()

	}
}
