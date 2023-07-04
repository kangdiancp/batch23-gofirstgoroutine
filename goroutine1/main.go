package main

import (
	"fmt"
	"time"
)

func printMe() {
	fmt.Println("Hello Go")
}

func main() {
	fmt.Println("Execute Goroutine")

	// child goroutine
	go printMe()

	// block main goroutine from terminating
	time.Sleep(time.Millisecond * 10)
	fmt.Println("Stop Goroutine")

}
