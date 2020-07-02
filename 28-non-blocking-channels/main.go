package main

import "fmt"

/*
	the default keyword inside a select statement just creates a non-blocking
	channel operation;

	if none of the cases have any messages available the default case is a non-
	blocking way out of the select

	it acts just like a switch statement, which it's what the select statement
	looks like it's trying to emulate the sintax of.
*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	// no message was put in the messages channel, therefore this will never
	// resolve
	case msg := <-messages:
		fmt.Println("recieved message", msg)

	// in this situation, the default case will resolve and "no message recieved"
	// will be printed
	default:
		fmt.Println("no message recieved")
	}
	// this also shows that the order is key in a select, whichever channel is
	// available will exit the select, even if other channels underneath it are
	// also available (default is always available)

	// in this case, the messages channel has no buffer, therefore the "hi" message
	// is lost because there's no imediate reciever
	msg := "hi"
	select {
	// this cannot resolve because the messages channel has no buffer
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// this is just a multi-case non blocking statement, similar to the other two
	// the default clause will be the one to execute as well
	select {
	case msg := <-messages:
		fmt.Println("recieved message", msg)
	case sig := <-signals:
		fmt.Println("recieved signal", sig)
	default:
		fmt.Println("no activity")
	}
}

/*
	output:
	no message recieved
	no message sent
	no output
*/
