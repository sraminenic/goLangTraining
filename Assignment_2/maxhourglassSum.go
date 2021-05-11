package main

import (
	"fmt"
)

func maxHourGlassSum(arr [6][6]int32) int32 {
	var maxhourglassSum, hourglassSum int32 = 0, 0
	for vLen := 1; vLen < len(arr)-1; vLen++ {
		for hLen := 1; hLen < len(arr[vLen])-1; hLen++ {
			hourglassSum = (arr[vLen-1][hLen-1] + arr[vLen-1][hLen] + arr[vLen-1][hLen+1] + arr[vLen][hLen] + arr[vLen+1][hLen-1] + arr[vLen+1][hLen] + arr[vLen+1][hLen+1])
			if hourglassSum > maxhourglassSum {
				maxhourglassSum = hourglassSum
			}
		}
	}
	return maxhourglassSum
}

func main() {
	sampleArray := [6][6]int32{{1, 2, 3, 0, 0, 0},
		{0, 1, 0, 5, 5, 6},
		{2, 1, 1, 9, 8, 7},
		{5, 9, 2, 4, 4, 0},
		{6, 5, 0, 2, 0, 0},
		{9, -9, 1, 2, 4, 0}}
	fmt.Println(sampleArray, maxHourGlassSum(sampleArray))
}
