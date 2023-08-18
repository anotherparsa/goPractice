package main

import "fmt"

func main() {
	arr := []int{1, 43, 64, 2, 42, 10, 192}
	size := len(arr)
	fmt.Println("before sorting:")
	showArray(arr, size)
	bubbleSort(arr, size)
	fmt.Println("after sorting")
	showArray(arr, size)
}

func bubbleSort(arr []int, size int) {
	isunsorted := true
	for i := 1; i <= size; i++ {
		if isunsorted {
			for j := 0; j < (size - i); j++ {
				if arr[j] > arr[j+1] {
					swapping(arr, j, j+1)
					isunsorted = true
				}
			}
		}
	}
}

func showArray(arr []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("")
}

func swapping(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
