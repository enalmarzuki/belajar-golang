package main

import "fmt"

func main() {
	// Cara 1
	var name string

	name = "Ummi Azizah"
	fmt.Println(name)

	name = "Enal Marzuki"
	fmt.Println(name)

	// Cara 2
	var name2 = "Ummi azizah 2"
	fmt.Println(name2)

	name2 = "Enal Marzuki 2"
	fmt.Println(name2)

	// Cara 3
	name3 := "Ummi azizah 3"
	fmt.Println(name3)

	name3 = "Enal Marzuki 3"
	fmt.Println(name3)

	// Multiple declaration
	var (
		firstName = "Enal"
		lastName  = "Marzuki"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
