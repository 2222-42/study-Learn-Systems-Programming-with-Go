package main

import (
	"fmt"
	"net"
	"os"
)

const path = "socketFile"

// 8.2.1
func main() {
	os.Remove(path)
	listener, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
	// Unixドメインソケットでは、Closeしないと、ソケットファイルが残り続ける
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Accept %v\n", conn.RemoteAddr())
}
