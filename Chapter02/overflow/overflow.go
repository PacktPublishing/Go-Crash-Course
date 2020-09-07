package main

import (
	"fmt"
)

func main() {
	var a uint8 = 0xFF
	var b uint8 = 0x01
	var c = a + b
	fmt.Println(a, b, c)

	var x int8 = 0x7F
	var y int8 = 0x01
	var z = x + y
	fmt.Println(x, y, z)
}
