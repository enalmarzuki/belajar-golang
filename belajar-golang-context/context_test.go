package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f")) // data is there
	fmt.Println(contextF.Value("c")) // get data from its parent
	fmt.Println(contextF.Value("b")) // the data you are looking for does not exist, because of a different parent
	fmt.Println(contextA.Value("b")) // can't retrieve data from the child
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
			time.Sleep(1 * time.Second) // slow simulation
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter : ", n)

		if n == 10 {
			break
		}
	}

	cancel() // send signal cancel to context

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter : ", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter : ", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine : ", runtime.NumGoroutine())
}
