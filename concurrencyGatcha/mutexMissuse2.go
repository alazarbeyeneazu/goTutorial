package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m1, m2 sync.Mutex
	done := make(chan bool)
	fmt.Println("STARTED ")
	go func() {
		m1.Lock()
		defer m1.Unlock()
		time.Sleep(1)
		m2.Lock()
		m2.Unlock()
		fmt.Println("SIGNAL")
		done <- true
	}()
	go func() {
		m2.Lock()
		defer m2.Unlock()
		time.Sleep(1)
		m1.Lock()
		m1.Unlock()
		fmt.Println("SIGNAL")
		done <- true
	}()
	<-done
	fmt.Println("Done")
	<-done
	fmt.Println("Done ")
}
