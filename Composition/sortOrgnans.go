package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}
type organs []Organ

func (l organs) Len() int {
	return len(l)
}
func (s organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct{ organs }
type ByWait struct{ organs }

func (b ByName) Less(i, j int) bool {
	return b.organs[i].Name < b.organs[j].Name
}
func (w ByWait) Less(i, j int) bool {
	return w.organs[i].Weight < w.organs[j].Weight
}
func main() {
	s := organs{Organ{Name: "Brain", Weight: 332}, {Name: "Heart", Weight: 34}}
	fmt.Println(s)
	sort.Sort(ByName{s})
	fmt.Println(s)
	sort.Sort(ByWait{s})
	fmt.Println(s)
}
