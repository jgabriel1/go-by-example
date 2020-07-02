package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5) // buffer of 5
	done := make(chan bool)   // no buffer blocking channel

	go func() {
		// this will run indefinitely
		for {
			// recieves a job into the variable j and also the information that
			// there are any other jobs in the message queue (more bool)
			j, more := <-jobs

			// whenever the previous line get's executed, i.e. if something is
			// recieved from the channel, this if statement gets checked:
			if more {
				// this runs if there are still messages left to be read in the
				// channel
				fmt.Println("recieved job", j)
			} else {
				// if j just recieved the last job left in the jobs channel, this
				// will execute
				fmt.Println("recieved all jobs")

				// until this line runs, the main thread is blocked
				done <- true
				break
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		// send jobs 1, 2 and 3 sequentially
		jobs <- j
		fmt.Println("sent job", j)
	}

	// closes the channel and no more messages can be sent through it, even though
	// it has still a buffer and is non-blocking
	close(jobs)
	fmt.Println("sent all jobs")

	// unblocks the main thread and finishes the program
	<-done
}
