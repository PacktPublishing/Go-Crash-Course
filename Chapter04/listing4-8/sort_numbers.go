package main

import "fmt"

func foo() {
	fmt.Println("test")
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter the numbers: ")
// 	numbersStr, _ := reader.ReadString('\n')
// 	numbersStrSlice := strings.Split(numbersStr, ",")

// 	numbers := make([]int, len(numbersStrSlice))
// 	for i, numberString := range numbersStrSlice {
// 		numbers[i], _ = strconv.Atoi(strings.TrimSpace(numberString))
// 	}

// 	for i, currentNum := range numbers[:len(numbers)-1] {
// 		smallestNumIdx := i
// 		smallestNum := currentNum

// 		for j := i; j < len(numbers); j++ {
// 			if numbers[j] < smallestNum {
// 				smallestNum = numbers[j]
// 				smallestNumIdx = j
// 			}
// 		}

// 		numbers[i] = smallestNum
// 		numbers[smallestNumIdx] = currentNum
// 	}

// 	sortedNumbersStrSlice := make([]string, len(numbers))
// 	for i, num := range numbers {
// 		sortedNumbersStrSlice[i] = fmt.Sprintf("%d", num)
// 	}

// 	fmt.Printf("sorted numbers: %s\n", strings.Join(sortedNumbersStrSlice, ", "))
// }
