package main

import "fmt"

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately,
// but the function call is not executed until the surrounding function returns.
func main() {
	// defer function callbacks are stacked so they are called in reverse order
	defer fmt.Println("world")
	defer fmt.Println("!")

	fmt.Println("hello")
}

// hello
// !
// world

// uzywane np do zamykania plikow
