package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format("2006/01/02/ 03:04:05 MST"))

	dr, err := time.ParseDuration("10m30s")
	if err != nil {
		panic(err)
	}
	fmt.Println(dr.String())

	d := time.Date(2017, time.August, 26, 11, 50, 30, 0, time.Local)
	fmt.Println(d.String())

	t2, err := time.Parse(time.Kitchen, "11:30PM")
	if err != nil {
		panic(err)
	}
	fmt.Println(t2.String())

	unixTime := time.Unix(1503673200, 0)
	fmt.Println(unixTime.String())

	fmt.Println(time.Now().Add(3 * time.Hour))

	fileInfo, err := os.Stat("./sampleTime.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v前\n", time.Now().Sub(fileInfo.ModTime()))

	fmt.Println(time.Now().Round(time.Hour))
}
