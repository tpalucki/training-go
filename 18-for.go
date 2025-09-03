package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// in and post statements are optional
	// semicolons are not needed
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// runs forever
	for {
	}
}
