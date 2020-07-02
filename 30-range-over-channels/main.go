package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	// queue <- "two"

	// closing the channel before the values are recieved???
	// yes, it can be closed and the values still be accessed afterwards
	close(queue)

	// iterates over all the messages in the queue in FIFO order
	for elem := range queue {
		fmt.Println(elem)
	}
	// output:
	// one
	// two
}
