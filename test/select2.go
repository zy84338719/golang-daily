package main

import (
	"fmt"
	"time"
)

func producer2(n int) <-chan int {
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

// consumer read data from in channel, print it, and print
// all proccess count in each second
func consumer2(in <-chan int) <-chan struct{} {
	finish := make(chan struct{})

	t := time.NewTicker(time.Millisecond * 500)
	processedCnt := 0

	go func() {
		defer func() {
			fmt.Println("worker exit")
			finish <- struct{}{}
			close(finish)
		}()

		// in for-select using ok to exit goroutine
		for {
			select {
			case x, ok := <-in:
				if !ok {
					return
				}
				fmt.Printf("Process %d\n", x)
				processedCnt++
			case <-t.C:
				fmt.Printf("Working, processedCnt = %d\n", processedCnt)
			}
		}
	}()

	return finish
}

func main() {
	out := producer2(3)
	finish := consumer2(out)

	// Wait consumer exit
	<-finish
	fmt.Println("main exit")
}
