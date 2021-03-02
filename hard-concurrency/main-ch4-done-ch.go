package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("work terminated")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// do something interesting
					fmt.Println(s)
				case <-done:
					return
				}

			}
		}()
		return terminated
	}

	done := make(chan interface{})
	strings := make(chan string)
	terminated := doWork(done, strings)
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("canceling doWork goroutine...")
		close(done)
	}()

	go func() {
		i := 0
		for {
			time.Sleep(1 * time.Second)
			strings <- fmt.Sprintf("%s - %d", "hello", i)
			i++
		}

	}()

	<-terminated
	fmt.Println("Done!!")

}
