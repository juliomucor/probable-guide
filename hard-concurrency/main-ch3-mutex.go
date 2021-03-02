package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	//  A Mutex provides a concurrent-safe way to express exclusive access to a shared resource
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		//    this is a critical section, a memory that needs to e shared between multiple concurrent proceses
		count++
		fmt.Printf("incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("decrementing: %d\n", count)

	}

	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			decrement()
		}()

	}

	wg.Wait()
	fmt.Printf("it is done: %d\n", count)
}
