package main

import "os"

func main() {
	os.Mkdir("setting", 0755)
	os.MkdirAll("setting/myapp/networksettings", 0755)
}
