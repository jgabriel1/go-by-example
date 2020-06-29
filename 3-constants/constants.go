package main

import (
	"fmt"
	"math"
)

// variables can be used by children scopes (like python and js)
// constants can't be redeclared in the same scope:
const s = "constant"

// * constants can only be used with fundamental types (int, bool, etc...)
// * contants need to be initialized on declaration (just like js)

func main() {
	// const s = "not-constant" in this case it would print // not-constant
	// children scope override parent scope ("normal" behavior)
	fmt.Println(s)

	// * constants are not typed on declaration, they recieve types when needed
	const n = 500000000

	// constants are computed with arbitrary precision:
	const d = 3e20 / n // numbers can be declared as scientific notation by default
	fmt.Println(d)     // it gets printed with scientific notation

	// it gets converted to regular int and is printed with a bunch of zeros
	fmt.Println(int64(d))
	// * types can be used as "functions" to convert to them (like in python)

	// this function takes float64 as argument, the constant n is then converted to be passed
	fmt.Println(math.Sin(n))
}
