package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the next prize"))
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the prize"))
}

func pangrams(inputText string) string {

	if len(inputText) < 26 {
		return "Not Pangrams"
	}
	inputTextInByte := []byte(strings.ToLower(inputText))
	m := make(map[byte]int)
	for _, byteValue := range inputTextInByte {
	if byteValue >= 'a' && byteValue <= 'z' {
			m[byteValue] = 1
		}		
	}
	
	if len(m) == 26 {
		return "Pangrams"
	} else {
		return "Not Pangrams"
	}
	
}
