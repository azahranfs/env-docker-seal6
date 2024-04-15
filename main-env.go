package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	massage := os.Getenv("HELLO_MASSAGE")
	if massage == "" {
		massage = "Hello, World! (no set env)"
	}
	fmt.Fprintf(w, massage)
}

func main() {
	http.HandleFunc("/", helloWorld)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running at http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
