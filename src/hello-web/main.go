package main

import (
    "fmt"
    "net/http"
)

var portNumber = ":8080"


func main() {
    // Define a handler function for the root endpoint
    http.HandleFunc("/about/", About)
    http.HandleFunc("/", Home)

    // Start the server and listen on port 8080
    fmt.Println(fmt.Sprintf("Starting server on port %s...", portNumber))
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
        fmt.Println("Error starting server:", err)
    }
}