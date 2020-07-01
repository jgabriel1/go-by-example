package main

import "fmt"

// range iterates over arrays/slices/maps

func main() {
	nums := []int{2, 3, 4} // slice (no number inside [])
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum) // 10

	// the first value in range is the index and the second is the value itself
	// it's basically javascript's map function arguments but in reverse order
	// [index], [value] := range [array]
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i) // 1
		}
	}

	// as expected, range applied to maps will return key, value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// [key], [value] := range [map]
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
		// a -> apple
		// b -> banana
	}

	// to iterate only over the keys
	for k := range kvs {
		fmt.Println("key:", k)
		// key: a
		// key: b
	}

	// when iterating over strings, range returns the UNICODE value for each
	// character (or rune)? still confused about what a rune is
	for i, c := range "go" {
		fmt.Println(i, c)
		// 0 103
		// 1 111
	}
}
