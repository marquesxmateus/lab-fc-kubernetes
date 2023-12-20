package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new server
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8081", nil)
	fmt.Println("Server is running on port 80")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}
