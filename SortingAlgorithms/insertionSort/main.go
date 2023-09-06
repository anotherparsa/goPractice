package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 3, 4, 2, 1, 8, 9, 7, 6}
	size := len(arr)
	fmt.Println("before insertion sorting: ")
	showArray(arr, size)
	fmt.Println("after insertion sorting: ")
	insertionSort(arr, size)
	showArray(arr, size)
}

func insertionSort(arr []int, size int) {
	for i := 1; i < size; i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			swaPPing(arr, j+1, j)
			j--
		}
		arr[j+1] = temp
	}
}

func showArray(arr []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println("")
}

func swaPPing(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
