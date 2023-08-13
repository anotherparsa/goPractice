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
	fmt.Println("enter the first number: ")
	placeHolder, _ := reader.ReadString('\n')
	placeHolder = strings.Replace(placeHolder, "\n", "", -1)
	Num1, _ := strconv.Atoi(placeHolder)
	fmt.Println("enter the second number: ")
	placeHolder, _ = reader.ReadString('\n')
	placeHolder = strings.Replace(placeHolder, "\n", "", -1)
	Num2, _ := strconv.Atoi(placeHolder)
	min := 0
	if Num1 > Num2 {
		min = Num2
	} else {
		min = Num1
	}
	GCD := 1
	for i := 1; i <= min; i++ {
		if (Num1%i == 0) && (Num2%i == 0) {
			GCD = i
		}
	}

	fmt.Printf("the GCD between %v and %v is %v \n", Num1, Num2, GCD)

}
