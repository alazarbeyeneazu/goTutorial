package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollar float32

func (d dollar) GetDolar() string {
	return fmt.Sprintf("$ %.2f", d)
}

type database map[string]dollar

func (db database) List(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s ", item, price.GetDolar())
	}
}
func (db database) Add(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	if _, ok := db[item]; ok {
		mgs := fmt.Sprintf("Duplicated item %s", item)
		http.Error(w, mgs, http.StatusBadRequest)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		mgs := fmt.Sprintf("Ivalide price format %s", price)
		http.Error(w, mgs, http.StatusBadRequest)
		return
	}
	db[item] = dollar(p)
	fmt.Fprintf(w, "we added %s to database", item)
}
func (db database) Update(w http.ResponseWriter, r *http.Request) {
	itemName := r.URL.Query().Get("item")
	newPrice := r.URL.Query().Get("price")
	if _, ok := db[itemName]; !ok {
		http.Error(w, "Item not found ", http.StatusBadRequest)
		return
	}
	p, err := strconv.ParseFloat(newPrice, 32)
	if err != nil {
		http.Error(w, "Cannot process the price ", http.StatusBadRequest)
		return
	}
	db[itemName] = dollar(p)
	fmt.Fprintf(w, "Updated Sucessfully To price of  %s", newPrice)
}

func (db database) GetOneItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		http.Error(w, " Item not found ", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "The price for %s is %s\n", item, db[item].GetDolar())
}
func (db database) Delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		http.Error(w, " Item not found ", http.StatusNotFound)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "Item %s deleted", item)
}

func main() {
	db := database{"shoes": 34, "shock": 34, "sendal": 345}
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/create", db.Add)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/item", db.GetOneItem)
	http.HandleFunc("/delete", db.Delete)
	fmt.Println("Servig website at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
