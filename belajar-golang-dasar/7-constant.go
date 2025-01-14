package main

import "fmt"

func main() {
	const firstName string = "Enal"
	const lastName = "Marzuki"

	// Error
	// firstName = "Ummi"
	// lastName = "Azizah"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Multiple Declaration
	const (
		firstName2 string = "Ummi"
		lastName2         = "Azizah"
	)

	fmt.Println(firstName2)
	fmt.Println(lastName2)

}
