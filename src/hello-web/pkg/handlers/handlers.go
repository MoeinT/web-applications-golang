package handlers

import (
    "fmt"
    "net/http"
    "hello-web/pkg/renders"
    "html/template"
)

func Home(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
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

    renders.RenderTemplate(w, data, tmpl)
}

func About(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
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

    renders.RenderTemplate(w, data, tmpl)

}

