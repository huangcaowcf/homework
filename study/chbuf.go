package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1
	ch <- 2
	s1, s2 := <-ch, <-ch
	//s2 := <-ch
	fmt.Println(s1)
	fmt.Println(s2)
}
