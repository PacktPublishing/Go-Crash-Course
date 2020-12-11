package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", v, i, v*i)
	}
}
