package main

import "fmt"

func main() {
	var a = 140 // 0b10001100
	var b = 170 // 0b10101010

	fmt.Println("140 &  170 =", a&b)
	fmt.Println("140 |  170 =", a|b)
	fmt.Println("140 ^  170 =", a^b)
	fmt.Println("140 &^ 170 =", a&^b)
	fmt.Println("140 << 2   =", a<<2)
	fmt.Println("140 >> 2   =", a>>2)
}
