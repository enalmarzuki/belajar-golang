package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Enal"
	names[1] = "Marzuki"
	names[2] = "Ummi"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// error = index 3 out of bounds
	// names[3] = "Azizah"

	// membuat array langsung
	var values [3]int = [3]int{
		90,
		80,
		// ingat jika index ke 3 tidak di deklarasi maka otomatis pakai default value yaitu 0 atau "" tergantung tipe data nya
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// function array
	var newArray = [...]int{
		90,
		80,
		25,
	}

	/*
	* jika kita tidak tahu ingin berapa panjang array nya
	* kita bisa menggunakan [...]int "titik tiga" dengan catatan
	* kita harus langsung mengisi array nya

	* namun jika dibuat seperti ini
	* var names [...]string
	* ini tidak bisa krn tidak langsung diisi
	 */

	fmt.Println(newArray)
	fmt.Println(newArray[0])
	fmt.Println(newArray[1])
	fmt.Println(newArray[2])

	fmt.Println(len(newArray))
}
