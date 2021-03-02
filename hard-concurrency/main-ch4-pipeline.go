package main

import (
	"fmt"
)

func main() {
	//  the very first stage of the pipeline
	//  we'll allways have some batch of data that you need to convert to a channel
	//  converts a discrete set of values into a stream of data on a channel
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)

			//range over the variadic slice (of any kind of data maybe) that was passed in and sends the slices' values on the channel it created
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}

			}
		}()
		return intStream
	}

	multiply := func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)

		go func() {
			defer close(multipliedStream)

			for i := range intStream {

				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)

		go func() {
			defer close(addedStream)

			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}

			}
		}()
		return addedStream
	}
	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for i := range pipeline {
		fmt.Println(i)
	}
}
