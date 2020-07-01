package main

import "fmt"

/*
	slices are like more powerful and useful arrays.
	append, for example, is not allowed for arrays but it is for slices
*/

func main() {

	// builtin make function initializes a slice:
	s := make([]string, 3)
	fmt.Println("emp:", s) // it prints as an empty array of 3 spaces

	// you can get and set values just like arrays (will not be repeating)
	// it also has len property
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	// s needs to be re-asigned when appended
	// append does NOT alter the state of the array in place
	// append is a function that returns a new slice
	fmt.Println("Pre append:", s)

	s = append(s, "d")
	s = append(s, "e", "f") // can append recursively
	fmt.Println("Post append:", s)

	// slices can also be copied
	// if the copy target array is not the same size, copy function only prints
	// to the array as much as it can, kinda like a string can be "printed into"
	// an array in C
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slices slicing operation is the same as python's
	l := s[2:5] // starts at 2 all the way up to 4, 5 excluded
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l) // all the way up to 4

	l = s[2:]
	fmt.Println("sl1:", l) // from 2 to the end

	// note that the declaration is almost the same, except for the lack of size
	// inside the brackets, which characterizes this as a slice initialization
	t := []string{"g", "h", "i"} // this is a slice, not an array
	fmt.Println("dcl:", t)

	// higher order slices are also allowed and the internal slices do not need
	// to have the same size (like in higher order arrays)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // this prints a non-matrix structure
}
