package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// convention says the "math/rand" package can be used as "rand" (as below)
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("My next favourite number is", rand.Int31n(12))
}
