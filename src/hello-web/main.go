package main

import (
    "fmt"
    "net/http"
    "html/template"
)

var portNumber = ":8080"

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

    RenderTemplate(w, "templates/home.page.html", data)
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

    RenderTemplate(w, "templates/about.page.html", data)

}

// RenderTemplate renders the specified template with the provided data
func RenderTemplate(w http.ResponseWriter, templateFile string, data interface{}) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Println("Error parsing the HTML file:", err)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		fmt.Println("Error executing template:", err)
	}
}

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