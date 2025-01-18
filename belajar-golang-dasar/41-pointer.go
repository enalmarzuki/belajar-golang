package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Pass By Value
	address1 := Address{"Makassar", "Sulawesi Selatan", "Indonesia"}
	address2 := address1 // copy value

	address2.City = "Barru"
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) // berubah menjadi Barru

	//  Pass By Reference
	/**
	* Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
	* Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
	 */

	newAddress1 := Address{"Makassar", "Sulawesi Selatan", "Indonesia"}
	newAddress2 := &newAddress1 // Pointer

	// Atau
	// var newAddress1 Address = Address{"Makassar", "Sulawesi Selatan", "Indonesia"}
	// var newAddress2 *Address = &newAddress1 // Pointer

	newAddress2.City = "Pare"
	fmt.Println(newAddress1) // Ikut berubah
	fmt.Println(newAddress2)
}
