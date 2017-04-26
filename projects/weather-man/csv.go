package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

// WriteCSV CSV에 기록하는 함수
func WriteCSV() {
	// Write CSV
	file, err := os.Create("tmp/gwanduke.csv")
	if err != nil {
		panic(err)
	}

	// CSV Writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))

	wr.Write([]string{"gwanduke", "success"})
	wr.Write([]string{"gwanduke2", "double success"})
	wr.Flush()
}

// ReadCSV 존재하는 CSV를 출력하는 함수
func ReadCSV() {
	// Read CSV
	file, _ := os.Open("public/test.csv")

	rdr := csv.NewReader(bufio.NewReader(file))

	rows, _ := rdr.ReadAll()

	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s ", rows[i][j])
		}
		fmt.Println()
	}
}
