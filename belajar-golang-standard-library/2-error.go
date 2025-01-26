package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation error")
	NotFoundError   = errors.New("Not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "enal" {
		return NotFoundError
	}

	// sukses
	return nil
}

func main() {
	err := GetById("enal")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
