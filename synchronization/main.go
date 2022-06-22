package main

import (
	"fmt"
	"sync"
)

func do() int {
	// var check = make(chan bool, 1)
	var check sync.Mutex
	var n int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// check <- true
			check.Lock()
			n++ //Data race
			check.Unlock()
			// <-check
			wg.Done()
		}()
	}
	wg.Wait()
	return int(n)

}
func main() {
	fmt.Println(do())
}
