package main

import (
	"fmt"
	"time"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			out = nil
			fmt.Println("producer exit")
		}()

		for i := 0; i < n; i++ {
			fmt.Printf("send %d\n", i)
			out <- i
			time.Sleep(time.Millisecond)
		}
	}()
	return out
}

// consumer only read data from in channel and print it
func consumer(in <-chan int) <-chan struct{} {
	finish := make(chan struct{})

	go func() {
		defer func() {
			fmt.Println("worker exit")
			finish <- struct{}{}
			close(finish)
		}()

		// Using for-range to exit goroutine
		// range has the ability to detect the close/end of a channel
		for x := range in {
			fmt.Printf("Process %d\n", x)
		}
	}()

	return finish
}

func main() {
	out := producer(100)
	finish := consumer(out)

	// Wait consumer exit
	<-finish
	fmt.Println("main exit")
}
