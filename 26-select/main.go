package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	// goroutine function that adds a string to c1
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	// the point is to show that this function and the previous start executing
	// one right after the other, and the total execution time is around 2 sec
	// since both timers start almost at the same time
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {

		// this syntax means that all of the cases are being awaited simultaneously
		// that's why both timers start at the same time, the select allows you
		// to await multiple goroutines
		select {
		case msg1 := <-c1:
			fmt.Println("recieved", msg1)
		case msg2 := <-c2:
			fmt.Println("recieved", msg2)
		}
	}

	/*
		* the reason why there is a for loop wraping the selection is that this
		select needs to be executed twice to print both values;

		* it awaits both coroutines at the same time, but exits when at least one
		of them yields a value;

		* since c1 "wins the race", it gets printed first and the select is exited;

		* after that, the for loop comes back around and awaits again for the select;

		* since c1 has no more values to yield, because the message was already
		printed, the only channel that has to be awaited is c2;

		* therefore, after the second loop, c2 gets printed, even though it was
		already initialized in the first loop;
	*/
}
