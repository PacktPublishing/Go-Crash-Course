package main

import "fmt"

func main() {
	fmt.Println("Uppercase comes before Lowercase so A < a")
	fmt.Println("APPLE", "<", "apple", "=", "APPLE" < "apple")
	fmt.Println("a", "<", "apple", "=", "a" < "apple")

	fmt.Println("Numbers come before Lowercase letters")
	fmt.Println("1", "<", "banana", "=", "1" < "banana")

	fmt.Println("Numbers come after Lowercase letters")
	fmt.Println("1", ">", "Banana", "=", "1" > "Banana")

	fmt.Println("if the strings are the same they are nor greater or lesser")
	fmt.Println("apple", ">", "apple", "=", "apple" > "apple")
	fmt.Println("apple", "<", "apple", "=", "apple" < "apple")

	fmt.Println("if they are the same they are both greater-equal and lesser-equal")
	fmt.Println("apple", ">=", "apple", "=", "apple" >= "apple")
	fmt.Println("apple", "<=", "apple", "=", "apple" <= "apple")

	fmt.Println("If they are not the same they work as greater and lesser")
	fmt.Println("Apple", ">=", "apple", "=", "Apple" >= "apple")
	fmt.Println("Apple", "<=", "apple", "=", "Apple" <= "apple")
}
