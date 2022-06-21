package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// this is going to return one per call
func getOne(i int) []byte {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Printf("Current Id [ %d ]", i)
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
	res, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Stoped %s", err.Error())
		os.Exit(-1)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {

		fmt.Fprintf(os.Stderr, "For %d Got status code %s", i, res.Status)
		return nil
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalide Body %s ", err.Error())
		os.Exit(-1)

	}

	return body
}
func main() {
	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		fails  int
		data   []byte
	)
	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		defer output.Close()
	}
	fmt.Fprint(output, "[")
	defer fmt.Println(output, "]")

	for i := 1; fails < 2; i++ {

		if data = getOne(i); data == nil {
			fmt.Printf(" Fails %d", fails)
			fails++
			continue
		}
		if cnt > 0 {
			fmt.Fprint(output, ",")
		}
		// wirte down to  stdout
		_, err := io.Copy(output, bytes.NewBuffer(data))

		if err != nil {
			fmt.Fprintf(os.Stderr, " it stoped %s \n", err)
			os.Exit(-1)
		}
		fails = 0
		cnt++
	}
	fmt.Println("Done")
}
