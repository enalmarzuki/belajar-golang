package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Jhon asd",
		MiddleName: "Doe",
		LastName:   "Vale",
	}

	encoder.Encode(customer)

	fmt.Println(customer)
}
