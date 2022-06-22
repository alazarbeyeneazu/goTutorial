package main

import (
	"fmt"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}
	for i := range chans {
		go func(i int, ch chan<- int) {
			time.Sleep(time.Duration(1) * time.Second)
			ch <- i
		}(i+1, chans[i])

	}
	for i := 0; i < 12; i++ {
		select {
		case m1 := <-chans[0]:
			fmt.Println("recived from ch1 ", m1)
		case m2 := <-chans[1]:
			fmt.Println("Recived from ch2", m2)
		}
	}
}
