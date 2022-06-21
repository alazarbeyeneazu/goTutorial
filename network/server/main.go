package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todos struct {
	UserID    int    `json:userId`
	Id        int    `json:id`
	Title     string `json:title`
	Completed bool   `json:completed`
}

var form = ` 
<h1> Todo #{{.Id}} </h1> 
<div> {{printf "User %d " .UserID}}</div> 
<div> {{prinft "%s (completed : %t)" ,.Title , .Completed}}</div> 
`

func handler(w http.ResponseWriter, r *http.Request) {
	const base = "https://jsonplaceholder.typicode.com/"
	res, err := http.Get(base + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var todo Todos
	if err = json.NewDecoder(res.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// err = json.Unmarshal(body, &todo)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// tmpl := template.New("form")
	// tmpl.Parse(form)
	// tmpl.Execute(w, todo)
	// fmt.Println(todo.Title)
	fmt.Fprintf(w, todo.Title)

}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
