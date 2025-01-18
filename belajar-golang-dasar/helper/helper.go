package helper

import "fmt"

/**
* Jika nama nya diawali dengan hurup besar, maka artinya bisa diakses dari package lain, jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain

* jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain tetapi masih bisa di akses dari package yang sama
 */

var version = "1.0.0"
var Application = "Golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Enal")
	fmt.Println(version)
}
