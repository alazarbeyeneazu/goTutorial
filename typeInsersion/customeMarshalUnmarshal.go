package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Name    string   `json:"name"`
	Entries []string `json:"entries"`
}

func main() {

	str := `{"name": "Alazar", "entries": ["classA", "classB"]}`

	// var res Response
	// json.Unmarshal([]byte(str), &res)
	// fmt.Println(res.Name)

	var res map[string]interface{}
	json.Unmarshal([]byte(str), &res)
	name, ok := res["name"].(string)
	if !ok {
		fmt.Println("no way")
	}
	fmt.Println(name)

}
