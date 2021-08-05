package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("UserID: %d\n", os.Getuid())
	fmt.Printf("GroupID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("SubGroupID: %v\n", groups)
}
