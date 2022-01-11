package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the numbers: ")
	numbersStr, _ := reader.ReadString('\n')
	numbersStrSlice := strings.Split(numbersStr, ",")

	numbers := make([]int, len(numbersStrSlice))
	for i, numberString := range numbersStrSlice {
		numbers[i], _ = strconv.Atoi(strings.TrimSpace(numberString))
	}

	sort.Ints(numbers) // sorts numbers in ascending order

	sortedNumbersStrSlice := make([]string, len(numbers))
	for i, num := range numbers {
		sortedNumbersStrSlice[i] = fmt.Sprintf("%d", num)
	}

	fmt.Printf("sorted numbers: %s\n", strings.Join(sortedNumbersStrSlice, ", "))
}
