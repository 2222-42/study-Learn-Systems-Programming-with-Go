package main

import (
	"fmt"
	"net"
	"time"
)

const address = "224.0.0.1:9999"
const interval = 10 * time.Second

func main() {
	fmt.Println("Start tick server at " + address)
	conn, err := net.Dial("udp", address)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	start := time.Now()
	wait := start.Truncate(interval).Add(interval).Sub(start)
	time.Sleep(wait)
	ticker := time.Tick(interval)
	for now := range ticker {
		if _, err := conn.Write([]byte(now.String())); err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Println("Tick: ", now.String())
	}
}
