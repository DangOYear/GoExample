package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	fmt.Println(v.X)

	p := &v
	p.X = 3
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
