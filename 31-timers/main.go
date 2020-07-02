package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	// a timer has a channel that will notify when the timer ends

	// this calling of C on timer1 indicates that the main thread will be blocked
	// until the timer ends
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// creates a new timer that waits for 1 second
	timer2 := time.NewTimer(time.Second)

	// this goroutine awaits on the timer second timer since the channel is being
	// called inside of it
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// this executes really fast after the goroutine is created, therefore stopping
	// the timer before it has enough time to fire
	stop2 := timer2.Stop()

	// since timer2 is now stopped, this evaluates to true
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// to double down on the fact that timer2 is never going to fire, wait for the
	// remaining time to give it a chance to do so
	time.Sleep(2 * time.Second)

	// spoiler alert: it doesn't (as expected)
}
