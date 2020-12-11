package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("specify one argument")
		return
	}
	v, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	s := strconv.Itoa(v)
	fmt.Printf("%d - %q\n", v, s)
}
