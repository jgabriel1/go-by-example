package main

import (
	"fmt"
	"time"
)

// this function recieves a channel as argument and inserts true into it when
// it's operation is complete
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)

	// by running the function as a goroutine, the done channel will recieve true
	// whenever the worker() function finishes executing
	go worker(done)

	// this blocks the main thread until  the goroutine is finished
	<-done
}

/*
	the main point is that, all the up untill the done message was called, the
	worker function hadn't done anything yet;

	the process only starts when it's needed, i guess
*/
