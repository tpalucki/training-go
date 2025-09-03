package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	structLiteral := &Vertex{X: 1} // Y:0 is implicit
	fmt.Println(structLiteral)
}
