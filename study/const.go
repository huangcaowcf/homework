package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	PI = 3.1415926
	E  = 2.0
	G  = 9.8
)

const (
	A = 10 * iota
	B
	C
)
const (
	Hello = "hello"
)

func main() {
	var n float32
	n = PI

	var t int
	var f float32
	t = 10
	f = float32(t) / 3
	fmt.Println(f * 3)
	fmt.Println(f)
	t = int(f) * 10
	fmt.Println(int(f * 3))
	fmt.Println(t, f)
	fmt.Println(A, B, C, n, Hello)

	var x int64
	x = 10240101
	var y int8
	y = int8(x)
	fmt.Println(x, y)

	var s string
	s = strconv.Itoa(t)
	z, err := strconv.Atoi("abc123")
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(z)
	fmt.Println(s)

	var h int64
	rand.Seed(time.Now().Unix())
	h = rand.Int63()
	var j = strconv.FormatInt(h, 17)
	fmt.Println(j)
}
