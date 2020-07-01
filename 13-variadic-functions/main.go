package main

import "fmt"

/*
	variadic functions are functions that take any number of arguments
	fmt.Println is an example of a variadic function
*/

// this indicates that the function recieves any number of ints
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// variadic functions can be called normally as such
	sum(1, 2)    // [1 2] 3
	sum(1, 2, 3) // [1 2 3] 6

	// or slices can be passed and destructured into them
	nums := []int{1, 2, 3, 4}
	
	// this passes 1, 2, 3, 4 as args to the sum function
	sum(nums...) // [1 2 3 4] 10
}
