package main

import (
	"fmt"
)

func main() {

	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer close(completed)
			defer fmt.Println("doWork exited")
			/*for {
				select {
				case s := <-strings:
					fmt.Println(s)
				}

			}*/
			for s := range strings {
				fmt.Println(s)

			}

		}()
		return completed

	}
	completed := doWork(nil)
	<-completed
	fmt.Println("end of work")
}
