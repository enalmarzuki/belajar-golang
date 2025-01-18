package main

import (
	"fmt"
)

func Random() any {
	return true
}

func main() {
	var result any = Random()
	// var resultString string = result.(string) // type assertions
	// fmt.Println(resultString)

	// var resultInt int = result.(int) // panic: interface conversion: interface {} is string, not int
	// fmt.Println(resultInt)

	// Cara 2
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
