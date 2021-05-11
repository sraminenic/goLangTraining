package main

import (
	"fmt"
	"strings"
)

func wordCount(wordStr string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(wordStr)
	for _, v := range words {
		wordMap[v]++
	}
	return wordMap
}

func main() {
	fmt.Println(wordCount("Hello Welcome to word count app"))
	fmt.Println(wordCount("count the repeated word by word count"))
}
