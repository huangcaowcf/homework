package main

import "fmt"

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		if x < 100 {
			x, y = y, x+y
		} else {
			break
		}
	}
	close(c)
}
func main() {
	c := make(chan int, 20)
	go fib(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
