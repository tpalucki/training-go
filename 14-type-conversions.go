package main

import (
	"fmt"
	"math"
)

// Type (value) converts value to Type
func main() {
	var x, y int = 3, 4
	// explicit conversion is required by Go between types
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
