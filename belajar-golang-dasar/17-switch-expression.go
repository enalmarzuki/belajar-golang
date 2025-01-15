package main

import "fmt"

func main() {
	name := "Enal"

	switch name {
	case "Enal":
		fmt.Println("Hello Enal")
	case "Ummi":
		fmt.Println("Hello Ummi")
	default:
		fmt.Println("Hi, Boleh Kenalan")
	}

	// Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	// Switch tanpa kondisi -- disarankan pakai if saja klo mau buat kondisi seperti ini
	name = "asdasdasd"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
