package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	clientPath := filepath.Join(os.TempDir(), "unixdomainsocket-client")
	if err := os.Remove(clientPath); err != nil {
		// ない場合、削除しようとしてエラーが起きるので、logに書き出すだけにする
		log.Println(err)
	}
	conn, err := net.ListenPacket("unixgram", clientPath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	unixServerAddr, err := net.ResolveUnixAddr("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
	if err != nil {
		panic(err)
	}
	var serverAddr net.Addr = unixServerAddr

	if _, err := conn.WriteTo([]byte("Hello from client"), serverAddr); err != nil {
		panic(err)
	}

	log.Println("Receiving from server")
	buffer := make([]byte, 1500)
	length, _, err := conn.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}
	log.Printf("Received: %s\n", string(buffer[:length]))
}
