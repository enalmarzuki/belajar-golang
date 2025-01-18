package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var newAddress1 Address = Address{"Makassar", "Sulawesi Selatan", "Indonesia"}
	var newAddress2 *Address = &newAddress1 // Pointer

	newAddress2.City = "Pare"
	fmt.Println(newAddress1) // Ikut berubah
	fmt.Println(newAddress2)

	// tanda and (&) tidak mengubah value variable yang lain yang mengacu pada sturct Address, tpi itu seperti membuat struct baru
	// newAddress2 = &Address{"Pangkep", "Sulawesi Selatan", "Indonesia"}
	// fmt.Println(newAddress1) // Tidak berubah menjadi pangkep
	// fmt.Println(newAddress2)

	// tanda bintang (*) akan mengubah semua value variable yang mengacu kepada struct Address
	*newAddress2 = Address{"Maros", "Sulawesi Selatan", "Indonesia"}
	fmt.Println(newAddress1)
	fmt.Println(newAddress2)
}
