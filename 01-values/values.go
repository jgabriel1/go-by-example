/*
	types:
		-string
		-integer
		-float
		-boolean
*/

package main

import "fmt"

func main() {
	// string concat (only way?)
	fmt.Println("go" + "lang")

	// ints and floats
	fmt.Println("1+1 =", 1+1)

	fmt.Println("7/3 =", 7/3)         // int/int = 2
	fmt.Println("7.0/3.0 =", 7.0/3.0) // float/float = 2.333...
	fmt.Println("7/3.0 =", 7/3.0)     // int/float = 2.333... (both converted
	fmt.Println("7.0/3 =", 7.0/3)     // float/int = 2.333...  to float)

	// booleans
	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
}
