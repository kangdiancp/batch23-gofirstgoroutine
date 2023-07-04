package main

import (
	"fmt"
	"time"
)

// comunicating by sharing memory, ga recommended
var finished bool

func sayHi() {
	fmt.Println("Hello Go")
	finished = true
}

func main() {
	go sayHi()
	for !finished {
		fmt.Println("Child goroutine is not finished")
		time.Sleep(time.Millisecond * 10)
	}

	fmt.Println("Child goroutine is finished")
}
