package main

import (
	"fmt"
	"net"
)

// 6.4
func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Printf("connected: %v\n", conn.RemoteAddr())
}
