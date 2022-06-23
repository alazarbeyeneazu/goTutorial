package main

import "fmt"

func abc() {
	panic("OMG")
}
func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recovered : ", p)
		}
	}()
	abc()
}
