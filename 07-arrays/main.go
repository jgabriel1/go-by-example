package main

import "fmt"

func main() {
	// Initializes an "empty array" with all default zero-values of int
	// which is 0
	var a [5]int
	fmt.Println("emp:", a) // [0 0 0 0 0] arrays are printed without commas

	// setting the value for a specific position in the array
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// len is a builtin function
	fmt.Println("len:", len(a))

	// declaring and initializing arrays
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", b)

	// second degree arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	/*
		[
			[0 1 2]
			[1 2 3]
		]
	*/
}
