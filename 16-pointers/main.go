package main

import "fmt"

// this function recieves an int as argument, which will be **passed by value**;
// this means that, at the moment zeroval is called, a copy of the value passed
// will be reassigned as zero, but the original value will remain the same,
// since the copy that was altered lives only in the internal socpe of the function
func zeroval(ival int) {
	ival = 0
}

// in this case, on the other hand, a pointer (reference to a value) is passed
// as argument to the function;
// whenever the function is called, the value at the address passed will be
// switched to zero, therefore altering the original value;
// this case is an example of **passing by REFERENCE**
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", 1) // 1

	zeroval(i)                 // nothing really happens here
	fmt.Println("zeroval:", i) // 1

	fmt.Println("pointer before:", &i) // current address of i in memory

	zeroptr(&i)               // & means I'm passing the address to i in memory as arg
	fmt.Println("zeroptr", i) // 0

	fmt.Println("pointer after", &i) // should print out the same address as before
}
