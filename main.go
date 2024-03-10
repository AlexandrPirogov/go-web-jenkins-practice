// main.go

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Jenkins! This is a simple Go web application.")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
