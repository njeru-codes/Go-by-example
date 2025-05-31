package main

import (
	"fmt"
	"sync"
)

// Go-routines  are like coperatively scheduled co-routines managed by goruntime

// to start a thread use the keyword   go

func sayHello(greetings string, wg *sync.WaitGroup) {
	defer wg.Done() //signal that this gouritne is done
	fmt.Println("program says", greetings)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1) //// we're going to wait for one goroutine

	go sayHello("hello world", &wg)

	wg.Wait() //wait for gouritine calls to be Done excuting

}

/**
COMMON PITFALLS
________________________
1. not synicng with main --> use sync.WaitGroup or channels to wait.
2. data races(Multiple goroutines accessing the same data) --> Use sync.Mutex, sync/atomic, or channels.


WHEN TO USE GO-routines
__________________________
Web servers				Handle each request in a separate goroutine
Parallel tasks			Download multiple files concurrently
Background work			Logging, metrics collection, cleanup
Performance				Speed up IO-bound or independent CPU-bound work
*/
