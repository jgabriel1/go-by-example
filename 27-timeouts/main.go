package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	// this simulates an operation that will take 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// in this case, since the operation takes 2 seconds and the timeout delay
	// is of 1 second, the timeout will "win the race" and exit the select block
	// with the timeout message
	select {
	// this operation takes 2 seconds, it's too slow
	case res := <-c1:
		fmt.Println(res)

	// this operation sets the time limit to be awaited, since this is only 1
	// second, it finishes first, exiting the select block
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// in this case, the timeout takes 3 seconds and the regular operation takes
	// only 2, so the result "wins" and gets printed exiting the select block
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

/*
	output:
		timeout 1
		result 2

*/
