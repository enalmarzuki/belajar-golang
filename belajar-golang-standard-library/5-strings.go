package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Enal Marzuki", "Enal"))
	fmt.Println(strings.Split("Enal Marzuki", " "))
	fmt.Println(strings.ToLower("Enal Marzuki"))
	fmt.Println(strings.ToUpper("Enal Marzuki"))
	fmt.Println(strings.Trim("     Enal Marzuki  ", " "))
	fmt.Println(strings.ReplaceAll("Enal Marzuki", "Enal", "Ummi"))
}
