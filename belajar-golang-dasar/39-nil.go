package main

import "fmt"

/*
	- function Contoh itu error karena
	- Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel

	func Contoh(name string) string {
		if name == "" {
			return nill
		} else {
			return name
		}
	}
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}

func main() {
	/*
		hasil code dibawah adalah empty string atau kosong

		data := NewMap("")
		fmt.Println(data["name"])

		karena function NewMap mengirikan "",
		kemudian kita memaksa untuk menampilkan "name" nya
		fmt.Println(data["name"])
	*/

	data := NewMap("Enal")
	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
