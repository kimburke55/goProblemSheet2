// An echo web application.
// https://ianmcloughlin.github.io 

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	//h1 header printing guessing game
	fmt.Fprintln(w, "<h1>guessing game!</h1>")
	
}

func main() {
	http.HandleFunc("/", requestHandler)
	//sends the information to port 8080
	http.ListenAndServe(":8080", nil)
}