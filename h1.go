// An echo web application.
// https://ianmcloughlin.github.io 

package main

import (
	"fmt"
	"net/http"
	
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>guessing game!</h1>")
	
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}