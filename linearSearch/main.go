package main

import "fmt"

func main() {
	array := []int{1, 43, 34, 234, 63, 234, 643, 212, 6765, 45, 21}
	size := len(array)
	fmt.Println(linearSearch(234, size, array))
	fmt.Println(linearSearch(6765, size, array))
	fmt.Println(linearSearch(3424, size, array))

}

func linearSearch(target int, size int, array []int) int {
	counter := 0
	for i := 0; i < size; i++ {
		counter++
		if array[i] == target {
			return counter
		}
	}
	return -1
}
