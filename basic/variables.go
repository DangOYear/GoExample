package main

import "fmt"

var c, python, java bool

var j, k int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var cpp, swift, kotlin = true, false, "yes"
	fmt.Println(j, k, cpp, swift, kotlin)

	/**
	只能在函数内使用
	*/
	l := 3
	fmt.Println(l)
}
