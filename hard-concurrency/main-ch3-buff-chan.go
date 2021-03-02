package main

import (
	"bytes"
	"fmt"
	"os"
	//	"time"
)

func main() {
	// in-memory buffer
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	//  changes the capacity to watch how the output changes
	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)
		//      exit before the channel is read
		defer fmt.Fprintln(&stdoutBuff, "Producer done")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending %v\n", i)
			intStream <- i
		}
	}()

	//	time.Sleep(1 * time.Second)
	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v\n", integer)
	}

}
