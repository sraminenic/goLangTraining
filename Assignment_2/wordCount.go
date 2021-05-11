package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

var x = `DOO WOP DOO WOP DOO WOP
SHOOBY DOOBY DOO
WOO WOO`

func main() {
	fmt.Println(wordCount(x))
}

func wordCount(wordStr string) map[string][]int {
	scanner := bufio.NewScanner(strings.NewReader(wordStr))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	wordMap := make(map[string][]int)
	lineNo := 1
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _, v := range words {
			wordMap[v] = append(wordMap[v], lineNo)
		}
		lineNo++
	}
	return wordMap
}
