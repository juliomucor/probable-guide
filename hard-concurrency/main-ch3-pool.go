package main

import (
	"fmt"
	"sync"
	"time"
)

//Concurrent-safe implementation of the object pool patern
func main() {

	//the pool pattern is a way to create and make available a fixed number, a pool, of things for use. It's commonly used to constrain the creation of things that are expensive
	myPool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new instance")
			return struct{}{}
		},
	}

	myPool.Get() // this data type can be safely used by multiple goroutines
	instance := myPool.Get()

	myPool.Put(instance) //put back to the pool, the pool maintain the reference so it will not be cleaned up by the garbage collector

	myPool.Get() // first check wether there are any available instances within the pool to return to the caller, and if not, call its New member variable to create one

	var numCalcsCreated int

	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	//  set the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
			time.Sleep(10 * time.Millisecond) //doing something
		}()

	}

	wg.Wait()
	fmt.Printf("the #Calcs is %d\n", numCalcsCreated)
}

//Another common situation where a Pool is useful is for warming a cache of preallocated objects
