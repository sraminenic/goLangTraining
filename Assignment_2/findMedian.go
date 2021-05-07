package main

import (
	"fmt"
	"sort"
)

func main() {
	sampleSlice := []int{7, 2, 3, 1, 9, 10}

	fmt.Println(findMedian(sampleSlice))

}

func findMedian(arr []int) int {
	sort.Ints(arr)
	fmt.Println(arr)
	middleIndex := int(len(arr) / 2)
	if len(arr)%2 == 0 {
		return (arr[middleIndex] + arr[middleIndex-1]) / 2
	} else {
		return arr[middleIndex]
	}
}
