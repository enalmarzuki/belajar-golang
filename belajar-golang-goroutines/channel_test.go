package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
======= Create Channel =======
*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	/*
		* Send data to channel
		channel <- "Jhon Wick"

		* Receive data from channel
		data := <-channel

		* send data directly as parameter
		fmt.Println(<-channel)

		* don't forget to close the channel
		defer close(channel)
	*/

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Jhon Wick"
		fmt.Println("has finished sending data to channel")
	}()

	data := <-channel
	fmt.Println("DATA :", data)
	time.Sleep(5 * time.Second)
}

/*
======= Channel as Parameter =======
*/
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Cristiano Ronaldo"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println("DATA :", data)
	time.Sleep(5 * time.Second)
}

/*
======= Channel In & Out =======
*/
// (channel chan<- string) => channel only for send data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Cristiano Ronaldo"
}

// (channel <-chan string) => channel only for receive data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("DATA :", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

/*
======= Buffered Channel =======
*/
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Valentino Rossi"
		channel <- "Jhon Wick"
		channel <- "Captain America"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Done")
}

/*
======= Range Channel =======
*/
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	// avoid this ways
	// data := <- channel
	// data := <- channel
	// data := <- channel
	// data := <- channel

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Looping " + strconv.Itoa(i)
		}
		close(channel)
	}()

	// use for range better
	for data := range channel {
		fmt.Println("Receive Data :", data)
	}

	fmt.Println("Has Done")
}

/*
======= Select Channel =======
*/
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

/*
======= Default Select Channel =======
*/
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		default:
			fmt.Println("Waiting...")
		}

		if counter == 2 {
			break
		}
	}
}
