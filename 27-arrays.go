package main

import "fmt"

func main() {
	var a [2]string // declare 2 items array of strings
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// output:
// Hello World
// [Hello World]
// [2 3 5 7 11 13]
