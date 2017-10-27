//https://stackoverflow.com/questions/15834278/serving-static-content-with-a-root-url-with-the-gorilla-toolkit
//http://www.alexedwards.net/blog/serving-static-sites-with-go
package main

import (
    "net/http"
	
)
func requestHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "<h1>guessing game!</h1>")
	
}
func main() {
    
	//http.Handle("/", http.FileServer(http.Dir("/file")))
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/file", requestHandler)
    http.ListenAndServe(":8080", nil)
}