package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const (
		min = 1
		max = 10
	)
	rand.Seed(time.Now().Unix())
	n := min + rand.Intn(max-min+1)
	fmt.Printf("I thought of a number between %d and %d... Guess it!\n", min, max)
loop:
	for {
		var guess int
		fmt.Print("> ")
		if _, err := fmt.Scanf("%d\n", &guess); err != nil {
			fmt.Println("Insert a valid number", err)
			continue
		}
		switch {
		case guess > n:
			fmt.Println("Too high")
		case guess < n:
			fmt.Println("Too low")
		default:
			break loop
		}
	}
	fmt.Println("That's it!")
}
