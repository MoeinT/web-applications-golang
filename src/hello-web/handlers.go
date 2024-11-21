package main

import (
    "fmt"
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    
    var message string
    
    if name != "" {
        message = fmt.Sprintf("Hello, %s!", name)
    } else {
        message = "Hello, World!"
    }

    // Define the data to pass to the template
    data := struct {
        Message string
    }{
        Message: message,
    }

    RenderTemplate(w, data, "templates/home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    
    var message string
    
    if name != "" {
        message = fmt.Sprintf("Hello, %s, this is the about page!", name)
    } else {
        message = "Hello, this is the about page!"
    }

    // Define the data to pass to the template
    data := struct {
        Message string
    }{
        Message: message,
    }

    RenderTemplate(w, data, "templates/about.page.html")

}

