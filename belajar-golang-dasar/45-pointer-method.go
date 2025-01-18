package main

import "fmt"

type Man struct {
	Name string
}

// tidak berubah menjadi Mr. Enal krn pass by value
// func (man Man) Married() {
// 	man.Name = "Mr. " + man.Name
// }

// Solusi nya menggunakan pointer *Man
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	enal := Man{"Enal"}
	enal.Married()

	fmt.Println(enal)
}
