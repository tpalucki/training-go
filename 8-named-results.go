package main

import (
	"fmt"
)

// you can name the return values
func split(sum int) (x, y int) {
	// these are declared as variables ad the top of the function so we use assignment
	x = sum * 4 / 9
	y = sum - x
	return // this is called 'naked' return, returns named argument
}

func main() {
	// declaration + assignment with :=
	x := "This is declaration and assignment"
	fmt.Println(x)

	// declaration
	var y string
	// assignment
	y = "This is just assignment"
	fmt.Println(y)

	// using function
	a, b := split(17)
	fmt.Println(a, b)
}
