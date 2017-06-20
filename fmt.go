package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("fmt.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(f, "hello")
	fmt.Fprintln(f, "worldln")
	s := "hellp"
	n := 5
	fmt.Fprintf(f, "string is %s %d", s, n)
	var i, j, sum int
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			sum = i * j
			fmt.Printf("%d * %d = %d  ", j, i, sum)
		}
		fmt.Printf("\n")
	}

	f.Close()
}
