package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s [pid]\n", os.Args[0])
		return
	}

	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		panic(err)
	}

	process.Signal(syscall.SIGINT)
	process.Kill()

	// cmdで機能したら、Process構造体が格納されるので、その変数経由でシグナルの送信が可能
	cmd := exec.Command("./sample")
	cmd.Start()
	cmd.Process.Signal(os.Interrupt)
}
