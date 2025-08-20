package main

import "fmt"

// unassigned variables - default zero values
func main() {
	var i int
	var f float64
	var b bool
	var s string

	// The zero value is:
	// 0 		for numeric types,
	// false 	for the boolean type, and
	// "" 		(the empty string) for strings.
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
