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
	fmt.Println("please enter the number")
	placeholder, _ := reader.ReadString('\n')
	placeholder = strings.Replace(placeholder, "\n", "", -1)
	number, _ := strconv.Atoi(placeholder)
	flag := false
	if number == 2 || number%2 == 0 {
		fmt.Println("it's a composite number")
	} else {
		for i := 2; i < number; i++ {
			if number%i == 0 {
				flag = true
			}
		}
		if flag {
			fmt.Println("this is a composite number")
		} else {
			fmt.Println("this is a prime number")
		}
	}

}
