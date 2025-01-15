package main

import "fmt"

func main() {
	name := "Ummi"

	if name == "Ummi" {
		fmt.Println("Hello Ummi")
	} else if name == "Unhy" {
		fmt.Println("Hello Unhy")
	} else {
		fmt.Println("Hi, Boleh Kenalan")
	}

	// Short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
