package main

import "fmt"

func main() {
	/*	var firstname, lastname string
		var age int
		firstname = "kudou"
		lastname = "shota"
		age = 21
	*/
	//all in one line ok
	var (
		firstName, lastName, age = "John", "Doe", 32
	)
	fmt.Println(firstName, lastName, age)
}
