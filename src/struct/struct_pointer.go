package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employees := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employees)
	//use pointer and change the value
	employeeCopy := &employees
	employeeCopy.FirstName = "David"
	fmt.Println(employees)
}
