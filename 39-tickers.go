/**
Timers are for when you want to do something once in the future
- tickers are for when you want to do something repeatedly at regular intervals.


*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(7 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		case <-done:
			fmt.Println("Done!")
			return
		}
	}
}
