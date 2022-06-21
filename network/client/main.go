package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Todos struct {
	Id        int    `json:id`
	Title     string `json:title`
	Completed bool   `json:completed`
}

const url = "https://jsonplaceholder.typicode.com"

func main() {
	res, err := http.Get(url + "/todos/1")
	if err != nil {
		log.Fatal("error ", err)
		os.Exit(-1)
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal("Error ", err)
			os.Exit(-1)
		}
		var todo Todos
		json.Unmarshal(body, &todo)
		fmt.Println(todo)
	}
}
