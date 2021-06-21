package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "www.google.co.jp:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if _, err := io.WriteString(conn, "Get / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"); err != nil {
		panic(err)
	}
	// 以下のを使うと、io.Copyから一向に抜けないので、要修正。
	//req, err := http.NewRequest("GET", "http://www.google.co.jp", nil)
	//if err != nil {
	//	panic(err)
	//}
	//if err := req.Write(conn); err != nil {
	//	panic(err)
	//}

	// net.Connがio.Readerインタフェースでもあることを利用して、サーバーから返ってきたレスポンスをio.Copy を使って画面に出力
	writer, err := io.Copy(os.Stdout, conn)
	println("end2")
	if err != nil {
		panic(err)
	}
	println("written %v\n", writer)
}
