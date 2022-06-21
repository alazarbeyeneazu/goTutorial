//read and replace text
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//read and replace
	if len(os.Args) < 3 {
		fmt.Println("No enough arguments")
		os.Exit(-1)
	}
	old, newv := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)
	var s []string

	for scan.Scan() {
		if scan.Text() == "" {
			break
		}
		s = strings.Split(scan.Text(), old)
		t := strings.Join(s, newv)
		fmt.Println(t)

	}

}
