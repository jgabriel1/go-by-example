package main

import (
	"errors"
	"fmt"
)

/*
	-> GO's approach to errors is very different than other languages' Exceptions

	* functions being able to return multiple values at once opens the possibility
	for error checking/handling in the function return itself

	* errors in go are like any other 1st class citizen and can be passed and
	handled as such

	* this approach also makes it easier to see which functions expect a possible
	error message
*/

// error type is a built in interface and by convention are always the last return
func f1(arg int) (int, error) {
	if arg == 42 {
		// -1 being a default return value for this function when errors happen
		// errors module create new error messages
		return -1, errors.New("can't work with 42")
	}

	// if nothing goes wrong, the result should be passed normally and nil (NULL)
	// should be passed instead of an error object
	return arg + 3, nil
}

// it's possible to create types that return errors by declaring the Error method
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// same implementation as f1 but now with a custom error type argError
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// this is iterating over a slice that goes from 7 to 42 (in this case 42 is included)
	for _, i := range []int{7, 42} {

		// it calls f1 for each value of the iteration [7, 42] using e for the
		// default possible error message

		// since f1 is being called "inside", it's error message is being
		// imediately checked
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	// this loop does exactly the same as the previous one but with the special
	// error construct declared before
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 failed:", r)
		}
	}

	// this will return an argError type as e
	_, e := f2(42)

	/* this is a type assertion */
	// ok evaluates to true if e is of type argError
	if ae, ok := e.(*argError); ok {
		// ae is then the root construct that had the Error method that was
		// thrown when f2 got an error catched
		fmt.Println(ae.arg)  // 42
		fmt.Println(ae.prob) // can't work with it
	}
}
