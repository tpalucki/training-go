package main

import (
	"fmt"
)

func main() {
	// we can only refer to exported names: "The ones with Capital letter"
	fmt.Println(math.pi)
	// should be replaced with 'math.Pi' - which is exported version
}
