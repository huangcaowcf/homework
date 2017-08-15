package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("a.txt")
	t, _ := os.OpenFile("b.txt", os.O_CREATE|os.O_RDWR, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	f.WriteString("hello\n")
	f.Seek(2, os.SEEK_SET)
	f.WriteString("##")
	f.Close()
	//t.WriteString("world\n")

	r := bufio.NewReader(t)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	t.Close()

	c, _ := os.Open("c.txt")
	s := bufio.NewReader(c)
	for {
		line, err := s.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
	c.Close()

}
