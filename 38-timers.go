/**
timers are used to schedule a one -time event in the future
Go provides time.Timer through the time package.
It’s different from time.Ticker, which repeats periodically.

--> You can stop a timer before it fires with Stop()
--> You can reuse a timer using Reset()
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("timer starteed")

	timer := time.NewTimer(3 * time.Second)

	<-timer.C // wait for the timer to fire
	fmt.Println("timer expired")
}
