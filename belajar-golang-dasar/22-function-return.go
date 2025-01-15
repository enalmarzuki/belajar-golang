package main

import "fmt"

// jika kita mau return data di function
// maka wajib menuliskan tipe data kembalian nya
func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Enal")
	fmt.Println(result)
}
