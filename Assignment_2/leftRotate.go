package main

import "fmt"

func main() {
	sampleSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sampleSlice, leftRotation(sampleSlice, 5))
}

func leftRotation(arr []int, d int) []int {
	return append(arr[d:cap(arr)], arr[0:d]...)
}
