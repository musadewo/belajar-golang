package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvData := "Muhammad,Sadewo,Wicaksono\n" +
		"Joko,Kurniawan,Khannedy\n" +
		"Budi,Joko,Budi"

	r := csv.NewReader(strings.NewReader(csvData))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}