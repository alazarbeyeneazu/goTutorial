package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b = ByteCounter(l)
	return len(p), nil
}

func main() {
	var c ByteCounter
	v1, _ := os.Open("text1.txt")

	v2 := &c
	n, _ := io.Copy(v2, v1)

	fmt.Println("Copied", n, "Bytes")
	fmt.Println("Copied", c, "Bytes")

}
