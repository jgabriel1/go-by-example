package main

import "fmt"

/* for is the only loop operator in go */
func main() {

	// regular while / increase loop format
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// OG for loop from C / Java
	for j := 7; j <= 9; j++ {
		fmt.Println(j) // 1 2 3 4 5 6 7 8 9
	}

	// while True loop (runs indefinitely)
	for {
		fmt.Println("loop") // 1x loop
		break
	}

	// continue kw can be used to next the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 { // mod operator exists
			continue
		}

		fmt.Println(n) // 1 3 5
	}
}
