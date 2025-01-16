package main

import "fmt"

func endAppRecover() {
	fmt.Println("End App")

	message := recover()
	fmt.Println("Terjadi panic ", message)
}

// recover function yang salah
// func runAppRecover(error bool) {
// 	defer endAppRecover()

// 	// recover yang salah
// 	// ini cara recover yang salah karena, jika kita sudah memanggil function panic maka code yang ada dibawah nya tidak akan dijalankan lagi
// 	if error {
// 		panic("Ups Error")
// 	}
// 	message := recover()
// 	fmt.Println("Terjadi Error ", message)
// }

// recover function yang benar
// itu recover nya di jalankan didalam function defer nya
func runAppRecover(error bool) {
	defer endAppRecover()

	if error {
		panic("Ups Error")
	}
}

func main() {
	/*
		* * Recover
				- adalah function yang bisa kita gunakan untuk menangkap data panic
				- dengan recover, proses panic akan terhenti, sehingga program akan tetap dijalankan
	*/

	// runAppRecover(false)
	// runAppRecover(true)

	runAppRecover(true)
	fmt.Println("Enal marzuki")
}
