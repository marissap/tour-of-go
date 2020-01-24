package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// add up the 2 numbers before current
	x := 0
	y := 1
	z := 0
	return func() int {
		z, x, y = x, y, x+y // create a circular pattern
		return z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}