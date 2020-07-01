package main

import "fmt"

func main() {
	var a = "initial" // type is automatically attributed when declared, it's
	fmt.Println(a)    // declaration is optional

	var b, c int = 1, 2 // multiple variables can be initialized at once
	fmt.Println(b, c)   // both are initialized as ints by default in case int is not passed.

	var d = true
	fmt.Println(d)

	var e int      // when declared but not initialized, the type is needed
	fmt.Println(e) // value 0 gets attributed when called in this case

	f := "apple"   // this removes the need for var and
	fmt.Println(f) // gives it the string default type
}
