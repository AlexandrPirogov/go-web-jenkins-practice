// main.go

package main

import (
	"fmt"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Jenkins! This is a simple Go web application.\n")
	fmt.Fprintf(w, "123 Added new line with CI/CD\n")
	fmt.Fprintf(w, "New line))))")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 8080...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
