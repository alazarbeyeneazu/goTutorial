package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++

	}
	fmt.Printf("\nthe file has %d unique words\n", len(words))
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})

	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	for _, v := range ss {
		fmt.Println(v.Key, " Appear ", v.Value, "times")
	}
}
