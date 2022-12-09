package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Index)

	//sarving on port 8080
	err := http.ListenAndServe(":8080", nil)

	//if error
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// top page
func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<html><body><p>Hello This Is a Golang World!!</p></body></html>")
}
