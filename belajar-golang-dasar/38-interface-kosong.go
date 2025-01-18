package main

import "fmt"

// any itu sama dengan interface{} "interface kosong"

func Ups() any {
	// return 1
	// return true
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
