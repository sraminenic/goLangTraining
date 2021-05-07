package main

import "fmt"

func main() {

	sampleSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(reverseSlice(sampleSlice))

}

func reverseSlice(arr []int) []int {
	for start, end := 0, len(arr)-1; start < end; start, end = start+1, end-1 {
		arr[start], arr[end] = arr[end], arr[start]
		fmt.Println(arr)
	}
	return arr
}
