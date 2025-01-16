package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	// defer logging() artinya di akhir function runApplication() maka function logging() akan di jalankan, mekipun ada error
	defer logging()
	fmt.Println("Run application")

	// jika ada sebuah error di atas code ini maka function logging() tidak akan di jalankan
	// logging()
}

func main() {
	/*
		* * Defer
				- adalah function yang bisa kita jadwalkan
					untuk dieksekusi setelah sebuah function
					selesai di eksekusi
				- defer function akan selalu di eksekusi
					walaupun terjadi error di function yang kita eksekusi
	*/

	runApplication()
}
