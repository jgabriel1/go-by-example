package main

import "fmt"

/*
	Channels are pipes that connect concurrent goroutines

	* you can send send values form a goroutine into a channel and that value is
	then accessible by other goroutines through that channel.

	* it's like a place where values are "posted" and shared between goroutines
*/

func main() {
	// make syntax can create a new channel
	//  in this case, that channel will be recieving a value of type string
	messages := make(chan string)

	go func() {
		// this arrow syntax indicates that the string "ping" is being sent into
		// the messages channel created earlier
		messages <- "ping"
	}()

	// this arrow syntax, on the other hand, which is different from the previous
	// one, indicates that the message is being sent from the messages channel
	// into the variable msg and is now in "synchronous land safely"
	msg := <-messages
	fmt.Println(msg) // ping
}
