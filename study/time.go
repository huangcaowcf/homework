package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		  timer := time.NewTicker(time.Second)
			cnt := 0
			for _ = range timer.C {
				cnt++
				if cnt > 4 {
					timer.Stop()
					return
				}
				fmt.Println("hello")
			}
	*/
	c := time.After(time.Second * 3)
	<-c
	fmt.Println("done")

}
