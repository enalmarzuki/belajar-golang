package main

import "fmt"

func main() {
	// alias
	type noKTP string

	var ktpEnal noKTP = "111111111"

	// jika variable sudah ada sebelumnya
	var contoh string = "22222222"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(ktpEnal)
	fmt.Println(contohKTP)
}
