package main

import "fmt"

func main() {
	var c bool
	//a = true
	//c = false
	c = ("hello" == "world")
	if c {
		fmt.Println("true")
	} else {
		fmt.Println(3 / 2.0)
	}
	var a, b int
	a = 2
	b = 1
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a = a + 3 //a += 3
	fmt.Println(a % b)
	b = a << 1
	var d = a * 2
	fmt.Println(b, d)
	//if a>=b{}
	//x:= a & b
	//if a > b && b >10{}
	//if a > b || b >10{}
	//s := "hello"
	//s += "world"
	//if a == b || a != b{}
}
