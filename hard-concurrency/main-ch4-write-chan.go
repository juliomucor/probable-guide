package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandStream closure exited")
			defer close(randStream)
			//      the for-select loop
			for {
				select { //do some work with channels
				case randStream <- rand.Int():
				case <-done: //loop infinitely until it is stopped
					return
				default:
					//            doing non-preemptable work: meaning that the resource cannot be taken away. An example is a printer.
				}
			}
		}()

		return randStream
	}
	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("print 3 random ints")

	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	close(done)

	//simulate ongoing work
	time.Sleep(1 * time.Second)

}
