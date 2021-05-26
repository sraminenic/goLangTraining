// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// CHALLENGE
//  Add error handling to the feet to meters program.
//
// EXPECTED OUTPUT
//  go run main.go hello
//    error: 'hello' is not a number.
//
//  go run main.go what
//    error: 'what' is not a number.
//
//  go run main.go 100
//    100 feet is 30.48 meters.
// ---------------------------------------------------------

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetsToConvert]`

func main() {
	meter, err := convertFeetToMeter(os.Args)
	if err == nil {
		fmt.Println(meter)
	} else {
		fmt.Println("Error", err)
	}
}

func convertFeetToMeter(osArgs []string) (string, error) {
	if len(osArgs) != 2 {
		return "", errors.New("Wrong arguments")
	}
	feet, err := strconv.ParseFloat(osArgs[1], 64)
	var meters float64
	var conversionMessage string
	if err == nil {
		meters = feet * 0.3048
		conversionMessage = fmt.Sprintf("%g feet is %g meters.", feet, meters)
	} else {
		err = errors.New(fmt.Sprintf("'%s' is not a number", osArgs[1]))
	}
	return conversionMessage, err

}
