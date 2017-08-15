package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
	j := 5
	for j < 8 {
		fmt.Println(j, " \n ")
		j++
	}

	l := 5
	for {
		l = l + 1
		fmt.Println(l)
		if l > 10 {
			break
		}
	}
	for c, arg := range os.Args {
		fmt.Println(c, arg)
	}

	s := "hello中文"
	for b, arg := range s {
		fmt.Printf("%d %c\n", b, arg)
	}

	var e, r, res int
	e = 1
	r = 1
	res = e + r
	fmt.Printf("%d %d ", e, r)
	sum := 0
	for sum < 100 {
		sum = e + r
		e = r
		r = sum
		res += sum
		//fmt.Printf("%d ", sum)
	}
	fmt.Printf("%d ", sum)
	fmt.Printf("%d ", res-sum)
}
