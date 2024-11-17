package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Define a handler function for the root endpoint
    http.HandleFunc("/", hander)

    // Start the server and listen on port 8080
    fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func hander(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    
    var message string
    
    if name != "" {
        message = fmt.Sprintf("Hello, %s!", name)
    } else {
        message = "Hello, World!"
    }

    // Write the response and check for errors
    if _, err := fmt.Fprintf(w, message); err != nil {
        fmt.Println("Error writing response:", err)
    }
}