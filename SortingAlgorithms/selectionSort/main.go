package main

import "fmt"

func main() {
	arr := []int{3, 6, 2, 7, 23, 634, 6534, 233}
	size := len(arr)
	fmt.Println("before sorting")
	showArray(arr, size)
	fmt.Println("after sorting")
	selectionSort(arr, size)
	showArray(arr, size)
}

func selectionSort(arr []int, size int) {
	for i := 0; i < size-1; i++ {
		minIndex := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			swapping(arr, i, minIndex)
		}
	}
}

func swapping(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func showArray(arr []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Printf("%v ", arr[i])
	}
	fmt.Println(" ")
}
