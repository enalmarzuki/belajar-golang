package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Enal", "Marzuki", "Ummi"})
	_ = writer.Write([]string{"Ummi", "Azizah", "Mukaddim"})
	_ = writer.Write([]string{"Unhy", "Meera", "Farida"})

	writer.Flush()
}
