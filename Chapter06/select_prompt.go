package main

import (
	"fmt"
	"time"
)

func main() {
	promptTimer := time.NewTimer(5 * time.Second)
	nameChan := make(chan string, 1)

	go func() {
		var name string
		fmt.Printf("Please enter your name: ")
		fmt.Scanf("%s", &name)

		nameChan <- name
	}()

	select {
	case name := <-nameChan:
		fmt.Printf("Hello, %s!\n", name)
	case <-promptTimer.C:
		fmt.Println("\nTimeout!")
	}
}
