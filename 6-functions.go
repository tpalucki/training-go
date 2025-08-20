package main

import (
	"fmt"
)

// function declaration
// note name of the arg first, then type, returned method type at the end
func add(x int, y int) int {
	return x + y
}

// you can omit types for params of the same type
func add2(x, y int) int {
	return x + y
}

// you can return different
func swap(x, y string) (string, string) {
	return y, x
}

// you can name the return values
func split(sum int) (x, y int) {
	// these are declared as variables ad the top
	x = sum * 4 / 9
	y = sum - x
	return // this is called 'naked' return, returns named argument
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))

	// multiple return values:
	a, b := swap("world!", "hello")
	fmt.Println(a, b)

	// named return values:
	fmt.Println(split(17))
}
