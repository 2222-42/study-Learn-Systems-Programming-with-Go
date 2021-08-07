package main

import (
	"github.com/creack/pty"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("./check")
	stdpty, stdtty, _ := pty.Open()
	defer stdtty.Close()
	cmd.Stdin = stdpty
	cmd.Stdout = stdpty
	errpty, errtty, _ := pty.Open()
	defer errtty.Close()
	cmd.Stderr = errtty
	go func() {
		io.Copy(os.Stdout, stdpty)
	}()

	go func() {
		io.Copy(os.Stderr, errpty)
	}()

	if err := cmd.Run(); err != nil {
		panic(err)
	}

}
