package main

import "fmt"

func main() {
	var a = true
	var b = false

	fmt.Println("true && false =", a && b)
	fmt.Println("true || false =", a || b)
	fmt.Println("! true        = ", !a)
}
