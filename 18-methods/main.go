package main

import "fmt"

// go supports the definition of methods asscociated with structs

/*
	passing by value x passing by reference:

	the reciever type of methods can carry either a copy of the struct (pass by
	value) or a pointer to a struct (pass by reference);

	passing by reference is ultimately better since the object won't have to be
	copied everytime it's passed onwards, making the process more efficient.

	also, if passing by value, any void type methods that modify the object won't
	have any direct effect on the original object passed to the function
*/

type rect struct {
	width, height int // struct properties of the same type can be grouped like in functions
}

// the reciever name should be descriptive of the struct that defines it's type,
// using "this" or "self" is not a good practice (i don't agree but will follow)

// func ([reciever] [type]) [method name]([...method args]) [return type] {...code}
// reciever is a pointer to a struct (passing by reference)
func (r *rect) area() int {
	return r.width * r.height
}

// reciever of the direct struct type (passing by value)
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}

	fmt.Println("area: ", r.area())  // 50
	fmt.Println("perim:", r.perim()) // 30

	// go automatically handles being passed by reference or value for all defined
	// methods;
	rp := &r
	fmt.Println("area: ", rp.area())  // 50
	fmt.Println("perim:", rp.perim()) // 30
	// always try to use pointers when calling methods since it's the more efficient
	// way, as discussed before;
}
