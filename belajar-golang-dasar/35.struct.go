package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var enal Customer
	fmt.Println(enal)

	enal.Name = "Enal"
	enal.Address = "Makassar"
	enal.Age = 27

	fmt.Println(enal)
	fmt.Println(enal.Name)
	fmt.Println(enal.Address)
	fmt.Println(enal.Age)

	// Cara 2
	ummi := Customer{
		Name:    "Ummi",
		Address: "Wono",
		Age:     27,
	}
	fmt.Println(ummi)

	// cara 3 = kekurangannya, ini harus sesuai dengan urutan property yang ada di strcut nya
	unhy := Customer{"Unhy", "Makassar", 31}
	fmt.Println(unhy)
}
