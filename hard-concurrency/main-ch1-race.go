package main

import (
	"fmt"
	"sync"
	"time"
)

// A race condition occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained

func main() {
	var memoryAccess sync.Mutex //not idiomatic go, though, the order of operations in this program is still nondeterministic
	var data int
	// a data race: two concurrent processes are attempting to access the same area of memory and the way they are accessing the memory is not atomic

	// critical section
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	time.Sleep(3 * time.Millisecond) // this is bad! will never be logically correct
	//The takeaway here is that you should always target logical correctnes

	// critical section
	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("the value is %v\n", "0")
	} else {
		// critical section
		fmt.Printf("the value is %v\n", data)

	}
	memoryAccess.Unlock()

	// there is a data race and the output of the program will be completely nondeterministic
	// the chances including: print the value but this = 1

}
