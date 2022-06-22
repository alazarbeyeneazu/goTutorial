package main

import (
	"fmt"
	"log"
	"net/http"
)

type NextId chan int

func (ch NextId) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> you have got %d </h1>", <-ch)
}
func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i

	}
}

//to
func main() {
	var nextId NextId = make(chan int)
	go counter(nextId)
	http.HandleFunc("/", nextId.handler)
	fmt.Println("Listening on the port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
