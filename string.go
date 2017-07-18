package main

import "fmt"

func Sum(s []string, c chan string) {
	var sum string
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to c
}

func main() {
	c := make(chan string)
	s := []string{"hello", "golang", "c++", "world"}
	go Sum(s[:len(s)/2], c)
	go Sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)

}
