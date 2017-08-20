package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(1000 * time.Millisecond).C
	boom := time.After(5000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("eat")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
