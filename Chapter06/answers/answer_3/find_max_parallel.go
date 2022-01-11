/*
We will use sync package's RWMutex here. According to the docs, a RWMutex
"is a reader/writer mutual exclusion lock. The lock can be held by an
arbitrary number of readers or a single writer. The zero value for a
RWMutex is an unlocked mutex."

We'll use the .RLock() method to lock the mutex for readers, and
.Lock() method to lock for writers.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var nums = make([]int, 0)
var rwMutex = sync.RWMutex{}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter comma-separated numbers: ")
	numbersString, _ := reader.ReadString('\n')
	numbersSliceStr := strings.Split(numbersString, ",")
	for _, numberStr := range numbersSliceStr {
		num, _ := strconv.Atoi(strings.TrimSpace(numberStr))
		nums = append(nums, num)
	}

	fmt.Print("Enter goroutines count: ")
	goroutinesCountStr, _ := reader.ReadString('\n')
	goroutinesCount, _ := strconv.Atoi(strings.TrimSpace(goroutinesCountStr))

	if goroutinesCount > len(nums) {
		goroutinesCount = len(nums)
	}

	batchSize := len(nums) / goroutinesCount

	maxSoFar := -1

	wg := sync.WaitGroup{}
	for i := 0; i < goroutinesCount; i++ {
		start := batchSize * i
		end := batchSize * (i + 1)
		if i == goroutinesCount-1 { // when contiguous numbers don't evenly split among goroutines, we make the last batch (unevenly) large enough to cover the remaining numbers
			end = len(nums)
		}
		wg.Add(1)
		go updateMaxInNums(start, end, &maxSoFar, &wg)
	}
	wg.Wait()

	fmt.Printf("Largest number: %d\n", *&maxSoFar)
}

func updateMaxInNums(start, end int, maxSoFar *int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i < end; i++ {
		rwMutex.RLock()
		maxSoFarVal := *maxSoFar
		rwMutex.RUnlock()
		if nums[i] > maxSoFarVal {
			rwMutex.Lock()
			*maxSoFar = nums[i]
			rwMutex.Unlock()
		}
	}
}
