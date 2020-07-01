package main

import "fmt"

// func [name]([param1][type], [param2][type], ...) [return type] {
//		[function code]...
// }
func plus(a int, b int) int {
	return a + b
}

// when multiple params have the same type, they can be grouped like this
// same as (a int, b int, c int)
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3) // reassigning res, don't need :=
	fmt.Println("1+2+3 =", res)
}
