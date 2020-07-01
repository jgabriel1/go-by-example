package main

import "fmt"

/*
	channels can be buffered, i.e. recieve more than one value before any one of
	the values is recieved by a reciever thread.
*/

func main() {

	// this syntax creates a channel that recieves up to 2 strings and stores
	// them until a reciever is ready to do so;
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	// this whole thing works as FIFO (first-in-first-out), which means the
	// channel will yield the messages in the same order they were recieved;
	fmt.Println(<-messages) // buffered
	fmt.Println(<-messages) // channel
}

/*
	* if the 2 is not passed in the make function, the compiler does not throw a
	descriptive error of what happened;

	* it does not just overwrite the first message, it just throws a weird error
	message
*/
