package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Create a new server
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
	fmt.Println("Server is running on port 80")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	msg := fmt.Sprintf("Hello %s, you are %s years old!\n", name, age)

	w.Write([]byte(fmt.Sprintf("<h1>%s</h1>", msg)))
}
