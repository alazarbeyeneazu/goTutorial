package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	done := make(chan bool)
	go func() {

		m.Lock()
		defer m.Unlock()
		fmt.Println("Locked")

	}()

	go func() {
		m.Lock()
		defer m.Unlock()
		fmt.Println("SIGNAL")
		done <- true
	}()
	<-done
	fmt.Println("DONE ")

}
