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
	http.ListenAndServe(":8080", nil)
}
