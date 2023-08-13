package main

import "fmt"

func main() {
	array := []int{1, 23, 54, 65, 72, 124, 456, 876, 999}
	size := len(array)
	fmt.Println(binarySearch(54, size, array))
	fmt.Println(binarySearch(456, size, array))
	fmt.Println(binarySearch(4566, size, array))
}

func binarySearch(target int, size int, array []int) int {
	start := 0
	end := size - 1
	for start <= end {
		middle := (start + end) / 2
		if array[middle] == target {
			return middle
		} else if array[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return -1
}
