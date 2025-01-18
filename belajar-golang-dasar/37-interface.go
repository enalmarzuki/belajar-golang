package main

import "fmt"

// Interface adalah sebuah tipe data abstract yang digunakan sebagai kontrak
// Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
// Sebuah interface berisikan definisi-definisi method
// Biasanya interface digunakan sebagai kontrak
type HasName interface {
	GetName() string
}

func sayHelloInterface(value HasName) {
	fmt.Println("Hello ", value.GetName())
}

// Cara implementasi interface HasName
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Enal"}
	sayHelloInterface(person)

	animal := Animal{Name: "Singa"}
	sayHelloInterface(animal)
}
