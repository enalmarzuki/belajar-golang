package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "enal,marzuki,ummi\n" +
		"ummi,azizah,mukaddim\n" +
		"unhy,meera,farida"

	reader := csv.NewReader(strings.NewReader(csvString))

	for { // ini artinya infinite looping
		record, err := reader.Read() // record itu perbaris
		if err == io.EOF {           // End Of File = ada di baris terakhir
			break
		}
		fmt.Println(record)
	}
}
