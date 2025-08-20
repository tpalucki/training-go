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

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))

	// btw assignment
	a, b := swap("world!", "hello")
	fmt.Println(a, b)
}
