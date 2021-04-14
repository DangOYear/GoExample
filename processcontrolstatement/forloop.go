package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//初始化语句和后置语句可以省略
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//"while"语句
	for sum < 2000 {
		sum += sum
	}
	fmt.Println(sum)

	//无限循环
	for {

	}
}
