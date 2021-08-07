package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"syscall"
)

func main() {
	fmt.Printf("process id: %d\n", os.Getpid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())

	sid, _ := unix.Getsid(os.Getpid())
	fmt.Fprintf(os.Stderr, "Group ID: %d, Session ID: %d\n", syscall.Getpgrp(), sid)
}
