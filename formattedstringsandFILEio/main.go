package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//cat
func main() {
	var lc, wc, cc int
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error %v", err)
			continue
		}

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()
			cc += len(s)
			lc++
			wc += len(strings.Fields(s))

		}
		fmt.Printf(" Number of words %7d Number of charactors %7d Number of line %7d FIle name %s \n", wc, cc, lc, fname)

		// data, err := ioutil.ReadAll(file)
		// if err != nil {
		// 	fmt.Fprint(os.Stderr, err)
		// 	continue
		// }
		// fmt.Printf("File %v  has size of %d Byte", fname, len(data))
		//copyt the file to stdout
		// if _, err := io.Copy(os.Stdout, file); err != nil {
		// 	fmt.Fprint(os.Stderr, err)
		// 	continue
		// }
		fmt.Println("")
		file.Close()

	}
}
