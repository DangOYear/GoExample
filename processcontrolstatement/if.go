package main

import (
	"fmt"
)

func main() {
	var x int = 4
	if x < 5 {
		fmt.Println("小于5")
	} else {
		fmt.Println("大于5")
	}

	//在判断前可以完成执行一个简单的语句
	if x = x + 4; x < 5 {
		fmt.Println("小于5")
	} else {
		fmt.Println("大于5")
	}

	if y := 5; y < 5 {
		fmt.Println("")
	}
}
