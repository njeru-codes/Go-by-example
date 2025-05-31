/**
Timeouts are used you don't want your program to wait forever
for something like a network request, slow computation, or message.

go has :
	1. time.After
	2. context.WithTimeout

time.After -->	This is useful when you’re waiting for something via a
				channel, but want to give up after a certain time.

context.WithTimeout --> This is the preferred way for managing timeouts, especially in APIs,
						web servers, and tasks that should be cancelable.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Timeout or cancel:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	doWork(ctx)
}
