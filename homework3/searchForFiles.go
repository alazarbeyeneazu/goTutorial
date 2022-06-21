package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// {
// 	"month": "1",
// 	"num": 1,
// 	"link": "",
// 	"year": "2006",
// 	"news": "",
// 	"safe_title": "Barrel - Part 1",
// 	"transcript": "[[A boy sits in a barrel which is floating in an ocean.]]\nBoy: I wonder where I'll float next?\n[[The barrel drifts into the distance. Nothing else can be seen.]]\n{{Alt: Don't we all.}}",
// 	"alt": "Don't we all.",
// 	"img": "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
// 	"title": "Barrel - Part 1",
// 	"day": "1"
// },
type Comics struct {
	Num        int    `json:num`
	Img        string `json:img`
	Day        string `json:day`
	Month      string `json:month`
	Year       string `json:year`
	Title      string `json:title`
	Trasncript string `json:transcript`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No file given to read ")
		os.Exit(-1)
	} else if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "No search quary is given ")
		os.Exit(-1)
	}
	fileName := os.Args[1]
	var (
		comicSlice []Comics
		terms      []string
		input      io.ReadCloser
		cnt        int
		err        error
	)
	input, err = os.Open(fileName)
	if err != nil {
		fmt.Fprint(os.Stderr, " Cannot open the file ", err.Error())
		os.Exit(-1)
	}

	err = json.NewDecoder(input).Decode(&comicSlice)
	fmt.Println("i have read ", len(comicSlice))
	for _, value := range os.Args[2:] {
		terms = append(terms, strings.ToLower(value))
	}
outer:
	for _, comics := range comicSlice {
		title := strings.ToLower(comics.Title)
		transcript := strings.ToLower(comics.Trasncript)
		for _, value := range terms {
			if !strings.Contains(title, value) && !strings.Contains(transcript, value) {
				continue outer
			}
		}
		// fmt.Printf("The Commic is at https://xkcd.com/%d/info.0.json / Day : %s /Month : %s /Year : %s \n", comics.Num, comics.Day, comics.Month, comics.Year)
		fmt.Printf("The Commic image at %s / Day : %s /Month : %s /Year : %s \n", comics.Img, comics.Day, comics.Month, comics.Year)
		cnt++
	}
	if err != nil {
		fmt.Fprint(os.Stderr, "Json conversion error ", err.Error())
		os.Exit(-1)
	}

}
