package main

import (
	"fmt"
	"time"
)

func main() {

	// make a channel to recieve requests and loop over to input rquests in the
	// channel until it's buffer is full
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	// by closing the channel, it cannot recieve any more messages
	close(requests)

	// this is a channel that limits the rate of which the messages will be recieved
	// from the requests channel (in this case, the requests are the "external
	// resource")
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// this recieving operation will have a limiter that, in this case will
		// block the thread until 200 miliseconds have passed, limiting the rate
		// of recieved messages
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// it's also possible to create another channel that will limit the total
	// requests per period (size of the burst of requests)
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
