package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

// 3.4.3
func main() {
	conn, err := net.Dial("tcp", "www.google.com:80")
	if err != nil {
		panic(err)
	}

	if _, err := conn.Write([]byte("GET / HTTP/1.0\r\nHost: www.google.com\r\n\r\n")); err != nil {
		panic(err)
	}
	// net.Conn型のconn(通信内容そのもの)をそのままコピーするのは効率的ではない。
	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Header)
	defer res.Body.Close()
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		panic(err)
	}
}
