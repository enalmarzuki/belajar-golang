package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You Are Blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Enal", blackList)
	registerUser("Anjing", blackList)

	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
