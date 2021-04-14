package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go runs on")
	var test string = "test"
	switch test {
	case "test":
		fmt.Println("ok")
	default:
		fmt.Println("not ok")
	}

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//没有条件的switch，类似于if else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Evening")
	}
}
