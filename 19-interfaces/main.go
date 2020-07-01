package main

import (
	"fmt"
	"math"
)

// geometry interface groups methods that have the same names
// it's an encapsulation of one or more methods
// * the types represent the return types of each method
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// implementation of geometry on rect type structs
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// implementation of geometry on circle type structs
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// interfaces are types therefore can be passed onto functions
func measure(g geometry) {
	fmt.Println(g)

	// methods passed as the interface can be called
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// if any of the objects passed as interfaces don't have the appropriate
	// methods defined by it, the linter will show an error and it will happen
	// at compile time as well
	measure(r)
	measure(c)
}
