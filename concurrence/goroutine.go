package main

import (
	"fmt"
	"time"
)

func say(s string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(s)
}

func main() {
	go say("world")
	say("hello")
}
