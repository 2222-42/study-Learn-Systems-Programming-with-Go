package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// // ファイルへの書き込み
	//file, err := os.Create("test.csv")
	//if err != nil {
	//	panic(err)
	//}

	// 標準出力
	file := os.Stdout
	writer := csv.NewWriter(file)
	records := [][]string{
		{"firstName", "lastName", "userName"},
		{"Sijuuni", "Kuroku", "2222-42"},
		{"Daioh", "Keishiki", "FDDaioh"},
	}
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatalln("error writing record to csv: ", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
