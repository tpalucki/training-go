package main

import (
	"fmt"
)

// function declaration
// note name of the arg first, then type, returned method type at the end
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
