package main

import (
	"fmt"

	"github.com/myuser/calc"
)

func main() {
	total := calc.Sum(3, 5)
	fmt.Println("Total:", total)
	fmt.Println("Version:", calc.Version)
}
