package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

// 8.2.2
func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-sample")
	if err := os.Remove(path); err != nil {
		// ない場合、削除しようとしてエラーが起きるので、logに書き出すだけにする
		log.Println(err)
	}
	listener, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Server is running at " + path)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
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
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World\n")),
			}
			if err := response.Write(conn); err != nil {
				log.Fatal(err)
			}
			if err := conn.Close(); err != nil {
				log.Fatal(err)
			}
		}()

	}

}
