package main

import (
	"encoding/json"
	"os"
)

// 2.4.7
func main() {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "    ")
	encoder.Encode(map[string]string{
		"example": "encoding/json",
		"hello":   "world",
	})
}
