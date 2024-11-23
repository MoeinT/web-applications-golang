package handlers

import (
    "fmt"
    "net/http"
    "hello-web/pkg/renders"
    "hello-web/pkg/configs"
)

// Define a Handler struct to hold appConfig
type App struct {
    AppConfig *configs.AppConfig
}


func (h *App) Home(w http.ResponseWriter, r *http.Request) {
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

    renders.RenderTemplate(w, data, "templates/home.page.html")
}

func (h *App) About(w http.ResponseWriter, r *http.Request) {
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

    renders.RenderTemplate(w, data, "templates/about.page.html")

}

