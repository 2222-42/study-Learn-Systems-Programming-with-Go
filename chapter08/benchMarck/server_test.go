package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// 8.4
func BenchmarkTCPServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:18888")
		if err != nil {
			panic(err)
		}
		request, err := http.NewRequest("GET", "http://localhost:18888", nil)
		if err != nil {
			panic(err)
		}

		if err := request.Write(conn); err != nil {
			panic(err)
		}

		if _, err = http.ReadResponse(bufio.NewReader(conn), request); err != nil {
			panic(err)
		}
	}
}

func BenchmarkUDSStreamServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		path := filepath.Join(os.TempDir(), "bench-unixdomainsocket-stream")
		conn, err := net.Dial("unix", path)
		if err != nil {
			panic(err)
		}
		request, err := http.NewRequest("get", "http://localhost:18888", nil)
		if err != nil {
			panic(err)
		}

		if err := request.Write(conn); err != nil {
			log.Println(err)
		}
		response, err := http.ReadResponse(bufio.NewReader(conn), request)
		if err != nil {
			panic(err)
		}

		if _, err := httputil.DumpResponse(response, true); err != nil {
			panic(err)
		}
	}
}

func TestMain(m *testing.M) {
	go UnixDomainSocketStreamServer()
	go TCPServer()
	time.Sleep(time.Second)
	code := m.Run()
	os.Exit(code)
}
