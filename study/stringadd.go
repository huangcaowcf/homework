package main

import "fmt"

func main() {
	s1 := "hello" + "world"
	s2 := "helloworld"

	if s1 == s2 {
		fmt.Println("equal")
	}
	fmt.Println(0, len(s1)-1)
	var c1 byte
	c1 = s1[0]
	fmt.Println(s1, s2, c1)
	fmt.Printf("%d %c\n", c1, c1)
	s3 := s1[:]

	fmt.Println(s3)
	var b byte
	for b = 0; b < 126; b++ {

		fmt.Printf("%d %c\n", b, b)
		fmt.Println(0xa)
	}

	array := []byte(s1)
	fmt.Println(array)
	array[0] = 72
	array[1] = 'E'
	s1 = string(array)
	fmt.Println(s1)
	fmt.Println('a' + ('H' - 'h'))
	fmt.Println(toupper("abcdefgAFD,SFDSF.&*()"))
}

func toupper(s string) string {

	array := []byte(s)
	for i := 0; i < len(s); i++ {
		if array[i] >= 'a' && array[i] <= 'z' {
			array[i] = array[i] - 32
		} else if array[i] >= 'A' && array[i] <= 'Z' {
			array[i] = array[i] + 32
		} else {
			array[i] = '!'
		}
	}

	for i := 0; i < len(s); i++ {
		if array[i] == '!' {
			fmt.Println("Please enter letters!No letters will be replaced with \"!\"")
			break
		}
	}
	var result string
	result = string(array)
	return result
}
