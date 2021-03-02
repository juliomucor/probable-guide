package main

//when working with concurrent code, there are a few different options for safe operation. We've gone over two of them:
//1. Synchronization primitives for sharing memory (e.g. sync.Mutex)
//2. Synchronization via communicating (e.g. channels)

/*other two:
3. Inmutable data, if a concurrent proccess wants modify it, it must create a copy
4. Data protected by confinement: ad-hoc (through a convention, whether it by set by the group you work within,...) and lexical (only read and only write channel)*/

//an example of lexical
import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer

		//    a copy to an internal buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}

		fmt.Println(buff.String())
	}
	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")

	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	//  will get a copy of data distinct from the one in the calling function
	inmutable := func(data string) string {
		data = fmt.Sprintf("%s%s", data, "-other")
		return data
	}

	//  dereferences the pointer from its memory address to the current value at that address
	no_inmutable := func(data *string) string {
		*data = fmt.Sprintf("%s%s", *data, "-other")
		return *data
	}
	wg.Wait()

	s := "a string"
	t := inmutable(s)
	fmt.Println(s, " ---- ", t)
	no_inmutable(&s)
	fmt.Println(s)

	//  efficient string concatenation
	var b strings.Builder
	b.Grow(32)
	for i, p := range []int{1, 2, 3, 5, 7} {
		fmt.Fprintf(&b, "%d:%d, ", i, p)
	}

	s1 := b.String()    // no copying
	s1 = s1[:b.Len()-2] // no copying, removes trailing ", "
	b.Reset()
	fmt.Println(s1)
}
