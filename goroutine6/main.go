package main

import "fmt"

func sayHello(ch chan bool) {
	fmt.Println("Send bool value to channel")
	ch <- true
}

func main() {
	ch := make(chan bool)

	go sayHello(ch)

	result := <-ch
	fmt.Println("Goroutine receive data from channel with value : ", result)
}
