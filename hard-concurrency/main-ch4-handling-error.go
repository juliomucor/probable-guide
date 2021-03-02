package main

import (
	"fmt"
	"net/http"
)

//She suggests me separate your concerns, in general, your concurrent processes should send their errors to another part of the program that has complete information about the state of my program
func main() {

	//  here we've coupled the potential result with the potential error
	//  this represents the complete set of possible outcomes created from the goroutine and allows our main goroutine to make decisions about what to do when errors occur
	type Result struct {
		Error    error
		Response *http.Response
	}

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}

				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://google.com", "bad_request"}

	results := checkStatus(done, urls...)

	for res := range results {
		//    here we are able to deal with errors coming out from the goroutine
		if res.Error != nil {
			fmt.Printf("error: %v\n", res.Error)
			continue
		}
		fmt.Printf("%v\n", res.Response.Status)
	}
}
