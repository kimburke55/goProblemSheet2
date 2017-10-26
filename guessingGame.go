// An echo web application.
// https://ianmcloughlin.github.io 

package main

import (
	"fmt"
	"net/http"
	
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "guessing game!")
	
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}