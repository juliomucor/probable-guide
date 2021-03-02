package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	sayHello := func() {
		defer wg.Done()
		fmt.Println("say hello")
	}

	salutation := "goodbye"

	wg.Add(2)
	//goroutines execute within the same address space they were created in
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	go sayHello() // fork

	for _, greeting := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println(greeting)
		}() // the loop will exit before the goroutines are begun. This means the 'salutation' variable falls out of scope but is transferred to the heap holding a reference to the last value
	}

	for _, greeting := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)

		go func(greeting string) {
			defer wg.Done()
			fmt.Println("proper", greeting)
		}(greeting) // the proper way, pass in the current iteration's variable to the closure.A copy of the string struct is made.
	}

	//this is the join point
	wg.Wait()
	fmt.Println(salutation) //changed to 'welcome'
}
