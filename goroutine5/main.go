package main

import "fmt"

/* func sayHello(ch chan bool){
	fmt.Println("Send bool value to channel")

} */

func main() {
	// declare channel unbuffred
	ch := make(chan int, 2)

	//send value
	ch <- 15
	fmt.Println(<-ch)
	ch <- 30
	fmt.Println(<-ch)
	//receiver
	//result := <-ch

}
