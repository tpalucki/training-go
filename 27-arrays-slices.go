package main

import "fmt"

func main() {
	// arrays
	var a [2]string // declare 2 items array of strings
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// array length is part of the type so cannot by resized

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// output:
	// [Hello World]
	// [2 3 5 7 11 13]
	// Hello World

	// Slices
	var s []int = primes[1:4] // slice from index 1 to 4 (excluding 4)
	fmt.Println(s)

	slicesAreLikeReferencesToArrays()

	sliceLiterals()

	sliceDefaults()

	sliceLengthAndCapacity()

	nilSlice()

	creatingSlicesWithMake()

	appendToSlice()
}

// Slices are like references to arrays (pointers to arrays, with additional index and length)
// A slice does not store any data, it just describes a section of an underlying array.
// Changing the elements of a slice modifies the corresponding elements of its underlying array.
// Other slices that share the same underlying array will see those changes.
func slicesAreLikeReferencesToArrays() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiterals() {
	// it's like array without
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	// below are quivalents
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	// these slice expressions are equivalent:
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]
}

func sliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func creatingSlicesWithMake() {
	// make building function is used to initialize slices
	// you can specify length = 5
	a := make([]int, 5)
	printSlice2("a", a)

	// You can also specify a capacity
	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func appendToSlice() {
	fmt.Println("Append to slice --------------")
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}
