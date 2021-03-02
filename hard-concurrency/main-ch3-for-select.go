package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	workCounter := 0

	fmt.Println("Blocking on read...")
	//    the select statement is the glue that binds channels together; like channels are the glue that binds goroutines locally together and also globally, at the intersection of two or more components in a system
	//    the select statements are one of the most crucial things
	//    if none of the channels are ready, the entire select statement blocks

loop:
	for {
		select {
		//    case statements aren't tested sequentially but considered simultaneosly to see if any of them are ready
		case <-c:
			fmt.Printf("Unblocked %v later\n", time.Since(start))
			break loop
			//      default clause in case we'd like to do something if all the channels we're selecting against are blocking
		default:
			fmt.Printf("In default after %v\n\n", time.Since(start))
		}
		//simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop\n", workCounter)

}
