package main

import "fmt"

func main() {
	var a = 10
	var b = 2

	fmt.Println("10 + 2 =", a+b)
	fmt.Println("10 - 2 =", a-b)
	fmt.Println("10 * 2 =", a*b)
	fmt.Println("10 / 2 =", a/b)
	fmt.Println("10 % 2 =", a%b)
	a++
	fmt.Println("10 ++  =", a)
	b--
	fmt.Println(" 2 --  =", b)
}
