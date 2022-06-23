package main

func main() {
	var ch = make(chan int)
	go func(ok bool) {
		if ok {
			ch <- 1
		}
	}(false)
	<-ch
}
