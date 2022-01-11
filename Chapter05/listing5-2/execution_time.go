package main

import (
	"fmt"
	"time"
)

func executeAndMeasureTime(f func(string), msg string) {
	start := time.Now()
	f(msg)
	end := time.Now()

	fmt.Println("\nTime taken (ns):", end.Sub(start).Nanoseconds())
}

func main() {
	executeAndMeasureTime(func(s string) {
		for _, c := range s {
			fmt.Printf(string(c))
		}
	}, "Hello")
	executeAndMeasureTime(func(s string) {
		fmt.Printf(s)
	}, "Howdy")
}
