package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func TCPServer() {
	// TCPではlnをcloseしなくてもよい
	ln, err := net.Listen("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}

			_, err = httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}

			response := http.Response{
				StatusCode: 200, ProtoMajor: 1, ProtoMinor: 0,
				Body: ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			if err := response.Write(conn); err != nil {
				log.Println(err)
			}
			conn.Close()
		}()

	}
}
