package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	// Create a new server
	http.HandleFunc("/", Hello)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/config", ConfigMapExample)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":8000", nil)
	fmt.Println("Server is running on port 80")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	msg := fmt.Sprintf("Hello %s, you are %s years old!\n", name, age)

	w.Write([]byte(fmt.Sprintf("<h1>%s</h1>", msg)))
}

func ConfigMapExample(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	fmt.Fprintf(w, "My Family: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s, Password: %s", user, password)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startedAt)

	if uptime.Seconds() < 10 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error - Uptime: %v", uptime.Seconds())))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Success - Uptime: %v", uptime.Seconds())))
	}
}
