package main

import "fmt"

func main() {
	counter := 0

	increament := func() {
		fmt.Println("Increment")
		counter++
	}

	increament()
	increament()
	increament()

	fmt.Println(counter)
}
