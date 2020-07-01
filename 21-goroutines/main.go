package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// this is a syncronous regular call of f
	f("direct")

	// this creates a new goroutine that executes separately (concurrently)
	go f("goroutine")

	// a function can also be defined and executed imediately as a goroutine
	// using this syntax
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) // this waits for 1 second
	fmt.Println("done")
}
