package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{}) //instantiate a new condition, takes in a type sync.Locker iface. This is what allows the Cond type to facilitate coordination
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock() // the main entering Wait and Unlock, we once again enter the critical section for the condition so we can modify data
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock() // exit the critical section
		c.Signal()   // here we let a goroutine waiting on the condition know that something has occurred
	}

	for i := 0; i < 10; i++ {
		c.L.Lock() // we lock the Locker in the main goroutine

		for len(queue) == 2 {
			c.Wait() // This is a blocking call and the goroutine will be suspended, upon entering Wait, Unlock is called on c.L -> thus allows the other goroutine does its action, after exiting Wait, Lock is called on c.L
		}

		fmt.Println("Adding to queue")

		queue = append(queue, struct{}{})

		go removeFromQueue(2 * time.Second)
		c.L.Unlock()
	}
}
