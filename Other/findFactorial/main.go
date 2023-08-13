package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the number")
	placeHolder, _ := reader.ReadString('\n')
	placeHolder = strings.Replace(placeHolder, "\n", "", -1)
	number, _ := strconv.Atoi(placeHolder)
	fmt.Println(findFactorial(number))

}

func findFactorial(number int) int {
	answer := 1
	for i := 1; i <= number; i++ {
		answer *= i
	}
	return answer
}
