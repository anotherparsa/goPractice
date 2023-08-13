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
	fmt.Println("please choose numerical option: \n1-binary to decimal \n2-decimal to binary ")
	placeHolder, _ := reader.ReadString('\n')
	placeHolder = strings.Replace(placeHolder, "\n", "", -1)
	selection, _ := strconv.Atoi(placeHolder)
	if selection == 1 {
		fmt.Println(bin2dec())

	} else if selection == 2 {
		fmt.Println(dec2bin())
	} else {
		fmt.Println("please choose between 1 and 2")
	}
}

func bin2dec() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter the number: ")
	placeHolder, _ := reader.ReadString('\n')
	placeHolder = strings.Replace(placeHolder, "\n", "", -1)
	number, _ := strconv.Atoi(placeHolder)
	power := 1
	answer := 0
	for number != 0 {
		lastDigit := number % 10
		answer = answer + power*lastDigit
		power *= 2
		number /= 10
	}
	return answer
}

func dec2bin() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the number: ")
	placeHolder, _ := reader.ReadString('\n')
	placeHolder = strings.Replace(placeHolder, "\n", "", -1)
	number, _ := strconv.Atoi(placeHolder)
	power := 1
	answer := 0
	for number != 0 {
		reminder := number % 2
		answer = answer + power*reminder
		power = power * 10
		number /= 2
	}
	return answer
}
