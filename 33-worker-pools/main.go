package main

import (
	"fmt"
	"time"
)

// this worker function encapsulates a unit that processes data recieved through
// a channel and outputs the results into another channel
func worker(id int, jobs <-chan int, results chan<- int) {

	// worker iterates over all available mess
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)

		// the sleep here is only adding an interval to simulate an expensive task
		time.Sleep(time.Second)

		fmt.Println("worker", id, "finished job", j)

		// when the worker is done doing the work, it outputs the result into the
		// results channel
		results <- j * 2
	}
}

func main() {
	start := time.Now()

	// the buffer size for the maximum amount of queued jobs
	// note that the amount of jobs is less than the maximum possible concurrent
	// jobs, otherwise, it would be impossible to buffer all them before executing
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// this first loop creates the 3 workers but they don't do anything yet, since
	// there are no jobs in queue
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// this creates 5 jobs to be put into the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// closing the jobs channel means that there are no more jobs available to
	// process, all of them have been queued
	close(jobs)

	// this loop now iterates recieving all the results and freeing the results
	// channel, ending all the remaining goroutines
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	fmt.Println("took:", time.Since(start)) // took: ~2sec

	// with all 3 goroutines created, it is expected that 2 of them will execute
	// 2 jobs and one of them will execute only 1 job
}
