package main

import (
	"fmt"
)

func main() {
	//  a goroutine that clearly owns a channel, manages its lifecicle
	chanOwner := func() <-chan int {
		//  encapsulated creation of the channel
		resultStream := make(chan int, 2)
		go func() {
			//    it's responsability of the owner to close the channel
			defer close(resultStream)
			defer fmt.Println("Done sending")
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		//    return a read-only channel
		return resultStream
	}

	resultStream := chanOwner()
	//  a consumer that clearly manages blocking and closing a channel
	for result := range resultStream {
		fmt.Printf("received %d\n", result)
	}
	//                blocked until the chan is closed
	fmt.Println("Done receiving")
}
