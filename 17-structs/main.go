package main

import "fmt"

// structs are typed collections of fields
// they're basically more definitions for more complex types
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func main() {

	// these create the struct directly
	// the arguments can be passed positionally
	fmt.Println(person{"Bob", 20}) // {Bob 20}

	// they can be passed as labeled args too
	fmt.Println(person{age: 30, name: "Alice"}) // {Alice 30}

	// when no args are passed to fields, they recieve the default zero-value
	// for their respective type
	fmt.Println(person{name: "Fred"}) // {Fred 0}

	// this syntax yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// it is idiomatic and best practice to encapsulate the struct creation operation
	// within a constructor function/method (newPerson)
	fmt.Println(newPerson("Jon")) // {Jon 42}

	s := person{"Sean", 50}

	// fields can be accessed with a dot
	fmt.Println(s.name) // Sean

	// dots can also be used with pointers to those structs
	sp := &s
	fmt.Println(sp.name) // Sean

	// structs are also perfectly mutable
	sp.age = 51
	fmt.Println(sp.age)
}
