package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Nama sudah benar", counter)
		counter++
	}

	fmt.Println("Selesai")

	// For dengan Statement
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke", counter2)
	}

	// Iterasi array, slice, atau map menggunakan FOR
	names := []string{"Enal", "Marzuki", "Ummi"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Nama ", i+1, names[i])
	}

	// Iterasi array, slice, atau map menggunakan FOR RANGE
	for index, name := range names {
		fmt.Println("Index ", index+1, ":", name)
	}

	// contoh index tidak dipakai, diganti menggunakan _
	for _, name := range names {
		fmt.Println(name)
	}
}
