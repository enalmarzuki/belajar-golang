package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	/*
	* jika kita buat seperti ini untuk return type nya
	* (firstName, middleName, lastName string)
	* maka kita tidak perlu lagi deklarasi kan variable nya, seperti ini
	* firstName := "Enal"

	* dan jika tipe data nya sama semua seperti ini
	* karena ini (firstName string, middleName string, lastName string)
	* maka kita bisa singkat menjadi
	* (firstName, middleName, lastName string)

	* karena ini (firstName, middleName, lastName string)
	* sudah langsung membuat variable nya jga sekaligus
	* seperti dibawah, kita tinggal ubah saja value nya
	 */

	firstName = "Enal"
	middleName = "Marzuki"
	lastName = "Ummi"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()

	fmt.Println(a, b, c)
}
