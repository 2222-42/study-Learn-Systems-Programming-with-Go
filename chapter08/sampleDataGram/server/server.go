package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-server")
	if err := os.Remove(path); err != nil {
		// ない場合、削除しようとしてエラーが起きるので、logに書き出すだけにする
		log.Println(err)
	}
	conn, err := net.ListenPacket("unixgram", path)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

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
