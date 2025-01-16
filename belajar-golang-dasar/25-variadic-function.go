package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10))

	// misal kita sudah memiliki slice terlebih dahulu
	numbers := []int{10, 10, 10, 10}

	// numbers... = ini otomatis mengubah slice menjadi variable arguments (varags)
	fmt.Println(sumAll(numbers...))
}
