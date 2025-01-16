package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}
}

func main() {
	/*
		* * Panic
				- adalah function yang bisa kita gunakan untuk menghentikan program
				- panic function biasanya dipanggil ketika terjadi panic pada saat program kita dijalankan
				- saat panic function dipanggil, program akan terhenti, namun defer function akan tetap dieksekusi
	*/

	runApp(false)
	runApp(true)
}
