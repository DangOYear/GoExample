package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}

/*
返回多个值
*/
func swap(x, y string) (string, string) {
	return y, x
}

/**
命名返回值
*/
func split(sum int) (x, y int) {
	x = sum / 10
	y = sum % 10
	return
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add1(1, 3))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	x, y := split(12)
	fmt.Println(x, y)
}
