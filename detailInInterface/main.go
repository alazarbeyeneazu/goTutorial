package main

import "fmt"

type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("Error %s %s", e.path, e.err)
}
func XYZ(i int) *error {
	return nil
}
func main() {
	var err *error = XYZ(1)
	if err != nil {
		fmt.Println((*err).Error())
	} else {
		fmt.Println("OK!")
	}
}
