package main

import "fmt"

func main() {
	firstName := "Enal"
	lastName := "Marzuki"

	fmt.Println("Hello '", firstName, lastName, "'")

	// dibanding dibuat seperti diatas mending pakai format
	// seperti ini
	fmt.Printf("Hello '%s %s' \n", firstName, lastName)
}
