package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Enal")
	data.PushBack("Marzuki")
	data.PushBack("Ummi")

	fmt.Println(data.Len())

	// data.Front() method Front untuk ambil data yang paling didepan
	var head *list.Element = data.Front()

	fmt.Println(head.Value) // Enal

	// head.Next() method Next untuk ambil data selanjutnya
	next := head.Next() // Marzuki
	fmt.Println(next.Value)

	next = next.Next() // Ummi
	fmt.Println(next.Value)

	// cara menggunakan looping
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
