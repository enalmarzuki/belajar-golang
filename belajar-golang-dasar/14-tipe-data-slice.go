package main

import "fmt"

func main() {
	names := [...]string{"Enal", "Marzuki", "Ummi", "Azizah", "Unhy", "Meera"}
	slice1 := names[4:6]

	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	// Function in Slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // sabtu, minggu
	fmt.Println(daySlice1)

	/*
	* ketika kita buat slice dari array dan mengupdate datanya seperti dibawah,
	* maka data di array nya juga akan ikut berubah
	 */
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	/*
	* ketika kita menggunakan fungsi "append"
	* maka slice nya akan mengecek terlebih dahalu ke array nya
	* apakah data "Libur baru" itu masih ditampung atau tidak
	* jika sudah tidak bisa tampung
	* maka slice akan membuat array baru. (bawaan dari go)

	* karena seperti yg kita tahu ketika membuat sebuah array
	* kita harus tentukan dlu panjang nya dan kita tidak boleh
	* mengisi array tersebut lebih dari panjang yg sudah di tentukan
	 */
	daySlice2 := append(daySlice1, "Libur baru")
	// daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu", "Libur Baru"}

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	/*
	*	karena daySlice2 sudah menggunakan array baru
	* maka ketika kita ubah datanya
	* itu tidak akan mempengaruhi array "days"
	* karena daySlice2 sudah menggunakan array baru
	 */
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice2)

	// Function Make
	// newSlice := make([]tipe array, kapasitas panjang slice, kapasitas array)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Enal"
	newSlice[1] = "Enal"

	// error harusnya menggunakan append, krn newSlice := make([]string, 2, 5)
	// kita sudah menentukan panjang dari slice nya 2, jadi hanya bisa 0 , 1
	// newSlice[2] = "Enal"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Enal Lagi")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Ummi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Notes hati" jangan sampai terbalik antara membuat slice dan array
	iniArray := [...]int{1, 2, 3, 4, 5} // atau [4]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
