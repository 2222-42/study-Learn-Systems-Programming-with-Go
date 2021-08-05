package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"
)

type FileLock struct {
	l  sync.Mutex
	fd int
}

func NewFileLock(filename string) *FileLock {
	if filename == "" {
		panic("filename needed")
	}
	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONLY, 0750)
	if err != nil {
		panic(err)
	}
	return &FileLock{fd: fd}
}

func (l *FileLock) Lock() {
	l.l.Lock()
	if err := syscall.Flock(l.fd, syscall.LOCK_EX); err != nil {
		panic(err)
	}
}

func (l *FileLock) Unlock() {
	if err := syscall.Flock(l.fd, syscall.LOCK_UN); err != nil {
		panic(err)
	}
	l.l.Unlock()
}

func main() {
	l := NewFileLock("lock.go")
	fmt.Println("try locking...")
	l.Lock()
	fmt.Println("locked!")
	time.Sleep(10 * time.Second)
	l.Unlock()
	fmt.Println("unlock")
}
