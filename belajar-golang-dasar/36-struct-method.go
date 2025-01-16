package main

import "fmt"

type Customer2 struct {
	Name, Address string
	Age           int
}

func (customer Customer2) sayHello(name string) string {
	return "Hello " + name + ", my name is " + customer.Name
}

func main() {

	ummi := Customer2{
		Name:    "Ummi",
		Address: "Wono",
		Age:     27,
	}

	fmt.Println(ummi.sayHello("Enal"))

}
