package main

import "fmt"

func main() {
	fmt.Println(stringConcatCount("abcd"))
	fmt.Println(stringConcatCount("abab"))
}

func stringConcatCount(inputText string) int {

	m := make(map[rune]int)
	for _, c := range inputText {
		m[c]++
	}
	return len(m)

}
