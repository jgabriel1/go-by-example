package main

import (
	"fmt"
	"sync"
	"time"
)

// a pointer to the workgroup is passed (instead of a copy)
func worker(id int, wg *sync.WaitGroup) {

	// whenever the function finishes executing, the waitgroup get's notified
	// that the work is done (finished)
	defer wg.Done()
	// defer executes after the function returns, so, after the last line of the
	// function block is executed, since, in this case, there's no return

	fmt.Printf("Worker %d starting\n", id)

	// sleep is just here to simulate an expensive task
	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// a single waitgroup is declared, but weirdly not instantiated or initialized
	// in any way
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// a counter associated with the waitgroup is incremented
		wg.Add(1)
		// this add operation adds to a counter of tasks the wait group is currently
		// awaiting, the done function called inside the worker function, on the
		// other hand, decreases that counter by 1 each time a worker is done
		// executing

		// and a worker is created right after associated with said waitgroup
		go worker(i, &wg)
	}

	// now that all the concurrent processes a encapsulated by the waitgroup, a
	// wait method is called on it to wait for all the processes to be done
	wg.Wait()
}
