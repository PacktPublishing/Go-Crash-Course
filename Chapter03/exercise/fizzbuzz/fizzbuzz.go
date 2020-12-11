package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i <= v; i++ {
		if i != 1 {
			fmt.Print("\n")
		}
		done := false
		if i%3 == 0 {
			done = true
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			done = true
			fmt.Print("Buzz")
		}
		if !done {
			fmt.Print(i)
		}

	}
}
