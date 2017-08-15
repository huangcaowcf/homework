package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{2, 7, 1, 6, 8, 9}
	for _, n := range s {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Println(n)
		}(n)
	}
	time.Sleep(5 * time.Second)
}
