package main

import (
	"fmt"
	"strconv"
)

func generator(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		ch <- i

	}
	// close(ch)
}
func filter(src chan int, des chan int, prim int) {
	for i := range des {
		if i%prim != 0 {
			des <- i
		}
	}
	// close(des)
}
func seilve(limit int) {
	var ch = make(chan int)
	go generator(limit, ch)
	for {
		pnum, ok := <-ch
		if !ok {
			break
		}
		ch1 := make(chan int)
		go filter(ch, ch1, limit)

		ch = ch1
		fmt.Printf(strconv.Itoa(pnum), " ")
	}
}
func main() {
	seilve(100)
}
