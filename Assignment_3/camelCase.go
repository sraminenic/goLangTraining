package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(camelCaseWordCount(""))
	camelCaseCount, camelCaseCountError := camelCaseWordCount("saveChangesInTheEditor")
	fmt.Println(camelCaseCount)
	fmt.Println(camelCaseCountError)

}

func camelCaseWordCount(inputText string) (int, error) {
	camelCaseCount := 1
	if len(inputText) == 0 {
		return 0, errors.New("Empty String passed, Can't find word count")
	}
	inputTextInByte := []byte(inputText)

	for _, byteValue := range inputTextInByte {
		if byteValue >= 'A' && byteValue <= 'Z' {
			camelCaseCount++
		}
	}
	return camelCaseCount, nil
}
