package main

import "fmt"

func add(numbers ...int) int {
	var sum int
	for _, nuber := range numbers {
		nbr := nuber
		sum += nbr
	}
	return sum
}

func main() {
	fmt.Println(add(1, 2, 3, 5, 6, 7, 7, 77))
	sum := []int{1, 443, 535, 335, 2, 2, 23, 23}
	sum = append(sum, sum...)
	fmt.Println(add(sum...))
}
