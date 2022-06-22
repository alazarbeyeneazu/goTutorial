package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("start ")
	const tiketRate = 2 * time.Second
	stoper := time.After(tiketRate * 5)
	tickers := time.NewTicker(tiketRate).C
loop:
	for {
		select {
		case <-tickers:
			fmt.Println("Tiker")
		case <-stoper:
			break loop

		}
	}
	log.Println("end ")
}
