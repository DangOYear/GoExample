package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := [6]int{2, 3, 4, 5, 6}
	fmt.Println(b)
}
