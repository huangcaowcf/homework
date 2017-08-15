package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "123"
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(n)
	}

	if n, err := strconv.Atoi(s); err == nil {
		fmt.Println(n)
	}

	switch s {
	case "1":
		fmt.Println("s=", "1")
	case "2":
		fmt.Println("s=", "2")
	default:
		fmt.Println("s=", "default")
	}

}
