package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("Enal", "Marzuki")
	sayHelloTo("Ummi", "Azizah")
	sayHelloTo("Unhy", "Meera")
}
