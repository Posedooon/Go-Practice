package main

import "fmt"

func main() {
	firstName := "John"
	//dont forget the & to get the address of the variable
	updateName(&firstName)
	fmt.Println(firstName)
}

/*
func updateName(name string) {
    name = "David"
}
	��not change the name "David"
	local variable is "David",
	but not change the name "John"
*/

func updateName(name *string) {
	*name = "David"
}

/*��change the name "David"
because the pointer is "John"*/
