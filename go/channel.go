package main

import (
	"fmt"
	"time"
)

func main() {
	// basicChannel()
	// channelBuffering()
	// channelSynchronization()
	// ticker()
}

func basicChannel() {
	messages := make(chan string)
	msg := ""

	// push onto the channel
	go func() { messages <- "ping" }()

	// before blocking
	fmt.Println("before:", msg)
	// reading from a channel is a blocking
	// read from the channel
	msg = <-messages
	// after blocking
	fmt.Println("after:", msg)
}

func channelBuffering() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"
	// messages <- "one too many"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// fmt.Println(<-messages)
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func channelSynchronization() {
	done := make(chan bool, 1)
	go worker(done)

	// this is an example of a blocking receive
	// when waiting for multiple you may prefer a waitgroup
	<-done
}

// func ticker() {
// 	timer := time.NewTimer(time.Second)
//
// 	count := 0
// 	// for count < 5 {
// 	// 	// block using the tick from the timer
// 	// 	<-timer.C
// 	// 	fmt.Println(count)
// 	// 	count++
// 	// }
//
// 	// select {
// 	// case <-timer.C:
// 	// 	fmt.Println(count)
// 	// 	count++
// 	// }
// 	for count != 5 {
// 		select {
// 		case <-timer.C:
// 			fmt.Println(count)
// 			count++
// 			// if count == 5 {
// 			// 	timer.Stop()
// 			// 	break
// 			// }
// 		}
// 		// fmt.Println(count)
// 		// count++
// 		// if count == 5 {
// 		// 	timer.Stop()
// 		// }
// 	}
// 	timer.Stop()
// 	fmt.Println("done")
// }
