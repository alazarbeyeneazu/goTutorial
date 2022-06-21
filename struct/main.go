package main

type Employee struct {
	Name  string
	Age   int
	Level string
	Boss  *Employee
}

func main() {
	// var e Employee
	// e.Name = "Alazar"
	// e.Age = 26
	// e.Level = 2
	// alazar := Employee{
	// 	Name:  "Alazar",
	// 	Age:   26,
	// 	Level: "master",
	// }
	employees := map[string]*Employee{}
	employees["boss"] = &Employee{Name: "Alazar", Age: 26}
	employees["abigiya"] = &Employee{Name: "Abigia", Age: 24, Boss: employees["boss"]}

	employees["abigiya"].Age++

}
