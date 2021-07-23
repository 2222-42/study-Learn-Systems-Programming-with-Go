package main

import "net"

// 8.2.1
func main() {
	conn, err := net.Dial("unix", "./server/socketFile")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
}
