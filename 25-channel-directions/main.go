package main

import "fmt"

/*
	in a function that recieves channels as arguments, you can specify that they
	can only recieve or send messages (direction) inside that function's scope,

	that adds another layer of type safety to the function

	with this you can lock the direction of a "information pipeline"
*/

// this syntax means that pings channel is only allowed to recieve messages
func ping(pings chan<- string, msg string) {
	pings <- msg
	// recieve := <-pings // this would throw a compile-time error
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message") // message is passed into pings
	pong(pings, pongs)            // message is passed from pings to pongs

	fmt.Println(<-pongs) // message is required by the main thread and called on print
}
