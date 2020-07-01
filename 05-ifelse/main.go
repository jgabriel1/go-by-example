package main

import "fmt"

func main() {

	// regular if/else control flow
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// isolated if statement
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// variables can be declared here, they're available in the entire block
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// * no parenthesis in statements but curly braces are needed
	/* THERE ARE NO TERNARY OPERATORS IN GO, IT HAS TO BE DONE LIKE THIS */
}
