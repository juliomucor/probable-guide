package main

import (
	"fmt"
	"sync"
)

func main() {
	//  an unbuffered channel is a buffered channel w a capacity of 0
	stringStream := make(chan string)
	//    because a goroutine was scheduled, there was no guarantee that it would run before the process exited
	go func() {
		//    this example works because channels in Go are said to be blocking
		stringStream <- "hello channels!" // the goroutine will not exit until the write is succesful
	}()
	// will sit there until a value is placed on the channel
	//  reads from a channel block if the channel is empty
	fmt.Println(<-stringStream)

	valueStream := make(chan interface{})
	close(valueStream)

	//  we could continue performing reads on this channel indefinitely despite the channel remaining closed
	// this opens up a few new patterns
	value, ok := <-valueStream
	fmt.Printf("using a closed channel  (%v): - %v\n", ok, value)

	//  ranging over a channel, range supports channels as arguments, and will automatically break the loop when a channel is closed.
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
	fmt.Println()
	//    unblocking multiple goroutines at once

	beginStream := make(chan interface{})
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			<-beginStream
			fmt.Printf("%v has begun\n", x)
		}(i)
	}

	fmt.Println("unblocking goroutines...")
	close(beginStream)
	wg.Wait()
}

//Thus, the main goroutine and the anonymous one block deterministically
