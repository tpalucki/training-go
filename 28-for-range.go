package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// The range form of the for loop iterates over a slice or map.
// When ranging over a slice, two values are returned for each iteration.
// The first is the index, and the second is a copy of the element at that index.
func main() {
	rangeIntroduced()

	skipIndexOrValueInRange()
}

func skipIndexOrValueInRange() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}

func rangeIntroduced() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
