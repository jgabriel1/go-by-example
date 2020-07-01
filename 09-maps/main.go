package main

import "fmt"

func main() {
	// maps are go's dictionaries(python) or objects(js)
	// Associative data structures

	// this creates a map where the [keys] are of type string and the values
	// are of type int
	m := make(map[string]int)

	// values can be set like regular python dicts
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) // map[k1:7 k2:13]

	// get values
	v1 := m["k1"]
	fmt.Println("v1:", v1) // 7

	// maps have lenght which is the number of k:v pairs
	fmt.Println("len:", len(m))

	// entries can be deleted by the void function delete
	delete(m, "k2")
	fmt.Println("map:", m) // map[k1:7]

	// getting values is actually a bit more complex than it seems
	// it can return a "tuple" (pythonist things...) to be destructured being
	// first value the one stored by the strucure and the second being if that
	// value exists or not

	vl, prs := m["k1"]
	fmt.Println("prs:", prs) // prints false because the key "k2" doesn't exist in m
	fmt.Println("val:", vl)  // prints 0 which is the default zero-value for int

	// it's possible to declare and initialize the map at once
	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	fmt.Println("map:", n) // map[bar:2 foo:1]
}
