package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, stop int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the starting number")
	placeholder, _ := reader.ReadString('\n')
	placeholder = strings.Replace(placeholder, "\n", "", -1)
	start, _ = strconv.Atoi(placeholder)
	fmt.Println("enter where to stop")
	placeholder, _ = reader.ReadString('\n')
	placeholder = strings.Replace(placeholder, "\n", "", -1)
	stop, _ = strconv.Atoi(placeholder)
	sum := 0
	for i := start; i <= stop; i++ {
		sum += i
	}
	fmt.Printf("the sum between %v and %v is %v\n", start, stop, sum)

}
