package main

import "errors"

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func caesarCipher(lengthInputText int, inputText string, noOfRotation int) (string, error) {
	if lengthInputText > 100 || lengthInputText < 1 {
		return "error 1", errors.New("not valid input length")
	}

	if noOfRotation > 100 || noOfRotation < 0 {
		return "error 2", errors.New("not valid rotation length")
	}
	noOfRotation %= 26

	if noOfRotation == 0 {
		return inputText, nil
	}
	caesarCipherText := []byte(inputText)

	for index, byteValue := range caesarCipherText {
		if byteValue >= 'a' && byteValue <= 'z' {
			lowerCaseAlphabetIndex := int((byteValue - 'a'))
			caesarCipherText[index] = lowerCaseAlphabet[(lowerCaseAlphabetIndex+noOfRotation)%26]
		} else if byteValue >= 'A' && byteValue <= 'Z' {
			upperCaseAlphabetIndex := int((byteValue - 'A'))
			caesarCipherText[index] = upperCaseAlphabet[(upperCaseAlphabetIndex+noOfRotation)%26]
		}
	}
	return string(caesarCipherText), nil
}

func main() {

	println(caesarCipher(26, "abcdefghijklmnopqrstuvwxyz", 1))
	println(caesarCipher(26, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 3))

	println(caesarCipher(38, "Always-Look-on-the-Bright-Side-of-Life", 5))

	//

}
