package main

import (
	"fmt"
	"sync"
)

func main() {
	// using waitgroup
	var wait sync.WaitGroup
	wait.Add(1)

	// using anonymous function
	go func() {
		fmt.Println("Child goroutine 1 Execute")
		wait.Done()
	}()

	wait.Wait()

	// multiple goroutine
	goRoutine := 5
	wait.Add(5)
	for i := 0; i < goRoutine; i++ {
		//create goroutine
		go func(index int) {
			fmt.Printf("Id : %d goroutine\n", index)
			wait.Done()
		}(i)
	}
	wait.Wait()

}
