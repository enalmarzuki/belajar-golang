package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

/*
* The function executed by the goroutine should be without return values
* because the return value can't be received by the goroutine
 */

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}
