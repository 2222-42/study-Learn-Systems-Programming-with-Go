package main

import (
	"fmt"
	"os"
	"time"
)

// 2.4.2
func main() {
	// os.Stdoutは画面への出力、os.Stderrは標準エラー出力
	os.Stdout.Write([]byte("os.Stdout example\n"))
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())
	fmt.Fprintf(os.Stdout, "%d: The value of %s is %f\n", 3, "Pi", 3.141592)
}
