package main

import "fmt"

/*
	IMPORTANT:
	go can return multiple values in the same function;
	this is often used in idiomatic go where you can return both the result and
	error values from a function
*/

// func [name]() ([return type 1], [return type 2]) ...
func vals() (int, int) {
	return 3, 7
}

func main() {
	// they can be destructured as expected
	// destructuring in go is called "multiple assignment"!!
	a, b := vals()
	fmt.Println(a) // 3
	fmt.Println(b) // 7

	_, c := vals()
	fmt.Println(c) // 7
}
