package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

// Deadlock
// a deadlocked program is one which all concurrent processes are waiting on one another. In this state, the program will never recover wo outside intervention
func main() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)

	//essentially, we have created two gears that cannot turn together
	//first call locks a and then attempts to lock b
	//but in the meantime, our second call has locked b and has attempted to locka
	go printSum(&a, &b)
	go printSum(&b, &a)

	wg.Wait()
}
