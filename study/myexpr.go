package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int
	//for i := 1; i < len(os.Args); i++ {
	//	fmt.Println(os.Args[i])
	//}
	if (len(os.Args) <= 3) || (len(os.Args) >= 5) {
		fmt.Println("Please enter a + b")
	} else if os.Args[2] == "+" {
		a, _ := strconv.Atoi(os.Args[1])
		b, _ := strconv.Atoi(os.Args[3])
		sum = a + b
		fmt.Println(sum)
	} else if os.Args[2] == "-" {
		a, _ := strconv.Atoi(os.Args[1])
		b, _ := strconv.Atoi(os.Args[3])
		sum = a - b
		fmt.Println(sum)
	}
}
