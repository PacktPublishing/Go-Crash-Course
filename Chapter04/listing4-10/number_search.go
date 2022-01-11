package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the numbers in increasing order: ")
	numbersStr, _ := reader.ReadString('\n')
	numbersStrSlice := strings.Split(numbersStr, ",")

	var num int
	fmt.Print("Enter the number to be searched: ")
	fmt.Scanf("%d", &num)

	numbers := make([]int, len(numbersStrSlice))
	for i, numberString := range numbersStrSlice {
		numbers[i], _ = strconv.Atoi(strings.TrimSpace(numberString))
	}

	low := 0
	high := len(numbers) - 1

	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if num > numbers[mid] {
			low = mid + 1
		} else if num < numbers[mid] {
			high = mid - 1
		} else {
			break
		}
	}

	if low > high { // number not found
		fmt.Printf("%d not found in provided list\n", num)
	} else {
		fmt.Printf("%d found at index %d in the provided list\n", num, mid)
	}
}
