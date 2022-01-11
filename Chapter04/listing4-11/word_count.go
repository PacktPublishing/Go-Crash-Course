package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("bunch_of_words.txt")

	text := strings.ToLower(string(data))
	reg, _ := regexp.Compile("[^a-z0-9 ]+")
	processedString := reg.ReplaceAllString(text, "")
	words := strings.Split(processedString, " ")

	wordCounts := make(map[string]int)
	for _, word := range words {
		if count, found := wordCounts[word]; found {
			wordCounts[word] = count + 1
		} else {
			wordCounts[word] = 1
		}
	}

	fmt.Println("Word counts:-")
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
