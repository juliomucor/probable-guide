package main

import (
	"fmt"
	"sync"
	"time"
)

//Starvation is any situation where a concurrent process cannot get all the resources it needs to perform work
// More broadly, starvation usually implies that one or more greedy concurrent process that are unfairly preventing one or more concurrent processes
//So starvation can cause your program to behave inefficiently or incorrectly.
func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	//Both workers do the same amount of simulated work(sleeping for 3 ns)
	greedyworker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("greedy worker: %v loops\n", count)
	}

	//constrain memory access synchronization only to critical sections
	//  performance(coarse-grained) vs fairness (fine grained)
	politeworker := func() {

		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("polite worker: %v loops\n", count)
	}

	wg.Add(2)
	go greedyworker()
	go politeworker()

	wg.Wait()

}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	time.Sleep(1 * time.Nanosecond)

	c.value++
}
