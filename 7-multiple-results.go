package main

import (
	"fmt"
)

// you can return different
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// multiple return values:
	a, b := swap("world!", "hello")
	fmt.Println(a, b)
}
