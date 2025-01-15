package main

import "fmt"

func getFullName() (string, string) {
	return "Enal", "Marzuki"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// menghiraukan salah satu return value nya
	firstName2, _ := getFullName()
	fmt.Println(firstName2)
}
