package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function for the root endpoint
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "" {
			fmt.Fprintf(w, "Hello, %s!", name)
		} else {
			fmt.Fprint(w, "Hello, World!")
		}
    })

    // Start the server and listen on port 8080
    fmt.Println("Starting server on port 8080...")
    
	err := http.ListenAndServe(":8080", nil)
	
	if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
