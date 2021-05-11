package main

import "fmt"

func main() {
	sampleSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sampleSlice, leftRotate(sampleSlice, 1))
}

func leftRotate(arr []int, d int) []int {
	if d > 0 && d <= len(arr) {
		return append(arr[d:cap(arr)], arr[0:d]...)
	} else {
		return arr
	}
}
