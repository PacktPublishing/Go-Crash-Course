package main

import (
	"fmt"
)

func main() {
	num := 0

	go func() {
		num = num*2 + 3
	}()

	go func() {
		num = num*3 + 2
	}()

	fmt.Println(num)
}
