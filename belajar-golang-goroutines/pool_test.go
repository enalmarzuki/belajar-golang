package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPoll(t *testing.T) {
	pool := sync.Pool{
		New: func() any { // override default value if the pool is empty
			return "New"
		},
	}

	pool.Put("Jhon")
	pool.Put("Doe")
	pool.Put("Alex")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("Data : ", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("DONE")
}
