package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

func UnixDomainSocketStreamServer() {
	path := filepath.Join(os.TempDir(), "bench-unixdomainsocket-stream")
	if err := os.Remove(path); err != nil {
		log.Println(err)
	}
	listener, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
	// Unixドメインソケットでは、Closeしないと、ソケットファイルが残り続ける
	defer listener.Close()

	for {
		conn, err := listener.Accept()
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
				StatusCode: 200,
				ProtoMinor: 0,
				ProtoMajor: 1,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			if err := response.Write(conn); err != nil {
				log.Println(err)
			}
			conn.Close()
		}()
	}
}
