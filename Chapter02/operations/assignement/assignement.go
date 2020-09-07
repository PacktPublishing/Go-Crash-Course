package main

import "fmt"

func main() {
	var a = 0

	var b = 2
	var c = 10

	fmt.Println("a =", a)
	a = b
	fmt.Println("a =", a)
	a = c
	fmt.Println("a =", a)

	a += b
	fmt.Println("a + 2 =", a)
	a -= b
	fmt.Println("a - 2 =", a)
	a *= b
	fmt.Println("a * 2 =", a)
	a /= b
	fmt.Println("a / 2 =", a)
	a %= b
	fmt.Println("a % 2 =", a)
	a &= b
	fmt.Println("a & 2 =", a)
	a |= b
	fmt.Println("a | 2 =", a)
}
