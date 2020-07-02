package main

import (
	"fmt"
	"time"
)

func main() {
	// tickers use a similar idea to timers but for periodic tasks, in this case,
	// the timer will be fired every .5 second
	ticker := time.NewTicker(500 * time.Millisecond)

	// blocking channel to make sure we stop the process when we're done
	done := make(chan bool)

	go func() {
		// indefinitely checks for channels done and ticker.C (channel associated
		// with the ticker), whenever the ticker channel "ticks", it prints the
		// time of when it ticked (that's what the channel associated with a ticker
		// returns)
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// stops the ticker after 1.6 seconds, so the expected amount of ticks is 3
	// before it stops
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()

	// throws a value into the done channel so that it exits the select statement
	// and returns the goroutine function, breaking the infinite loop of checking
	done <- true
	fmt.Println("Ticker stopped")
}
