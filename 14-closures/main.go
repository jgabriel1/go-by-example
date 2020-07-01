package main

import "fmt"

func intSeq() func() int {
	// i is an internal variable of this function's scope
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// when intSeq gets called, a new value of i is "instantiated"
	nextInt := intSeq() // nextInt is now a function with a value of i associated with it

	// whenever nextInt gets called, it's value of i get's increased
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// to prove that the value of i is attached to calling the intSeq function
	newInts := intSeq()    // i in this closure is now 0
	fmt.Println(newInts()) // 1

	// back to the first call
	fmt.Println(nextInt()) // 4
}

// the best way to understand closures is that a function remembers where it was created
