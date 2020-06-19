package main

import (
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}
	//io.WriteString(conn, "Get / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
	req, err := http.NewRequest("GET", "http://ascii.jp", nil)
	req.Write(conn)
	// net.Connがio.Readerインタフェースでもあることを利用して、サーバーから返ってきたレスポンスをio.Copy を使って画面に出力
	io.Copy(os.Stdout, conn)
}
