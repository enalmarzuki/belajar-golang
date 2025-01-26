package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	// Cara Manual menambahkan data di ring
	// data.Value = "Value 1"

	// data = data.Next()
	// data.Value = "Value 2"

	// data = data.Next()
	// data.Value = "Value 3"

	// data = data.Next()
	// data.Value = "Value 4"

	// data = data.Next()
	// data.Value = "Value 5"

	// Cara menggunakan looping untuk tambah data di ring
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i+1)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
