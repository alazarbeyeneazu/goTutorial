package main

import "fmt"

func do(b *map[int]int) {
	(*b)[3] = 0

	*b = make(map[int]int)
	(*b)[0] = 3

}
func main() {
	a := map[int]int{2: 1, 4: 2, 3: 3}
	fmt.Println(a)
	do(&a)
	fmt.Printf("\n%v\n", a)

}
