package main

import (
	"fmt"
	"net"
)

const address = "localhost:8888"

func main() {
	fmt.Println("Server is running at " + address)
	conn, err := net.ListenPacket("udp", address)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v", remoteAddress, string(buffer[:length]))
		if _, err := conn.WriteTo([]byte("Hello from Server"), remoteAddress); err != nil {
			panic(err)
		}
	}
}
