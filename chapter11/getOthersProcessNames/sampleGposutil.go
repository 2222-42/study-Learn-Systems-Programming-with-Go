package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os"
)

func main() {
	p, _ := process.NewProcess(int32(os.Getppid()))
	name, _ := p.Name()
	cmd, _ := p.Cmdline()
	fmt.Printf("parent pid: %d, name: '%s', cmd: '%s'\n", p.Pid, name, cmd)
}
