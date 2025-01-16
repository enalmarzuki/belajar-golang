package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello ", filteredName)
// }

// Function type declarations
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Enal", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

}
