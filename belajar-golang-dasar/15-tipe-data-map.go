package main

import "fmt"

func main() {
	// Cara 1
	// var person map[string]string = map[string]string{}
	// person["name"] = "Enal"
	// person["address"] = "Makassar"

	// Cara 2
	person := map[string]string{
		"name":    "Enal",
		"address": "Makassar",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	// karena key "salah" tidak ada di tipe data map
	// maka otomatis menggunakan default value sesuai tipe data nya
	// disini string maka default value nya adalah string kosong ""
	fmt.Println(person["salah"])

	// Function di tipe data Map
	book := make(map[string]string)
	book["title"] = "Buku golang"
	book["author"] = "Enal"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
