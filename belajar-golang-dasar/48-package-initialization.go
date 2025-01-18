package main

/**
* Blank Identifier :
	* Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
	* Secara default, Go-Lang akan komplen ketika ada package yang di import namun tidak digunakan
	* Untuk menangani hal tersebut, kita bisa menggunakan blank identifier (_) sebelum nama package ketika melakukan import
*/

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
