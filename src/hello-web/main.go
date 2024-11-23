package main

import (
    "fmt"
    "net/http"
    "hello-web/pkg/handlers"
    "hello-web/pkg/configs"
    "hello-web/pkg/renders"
    "log"
)

var portNumber = ":8080"


func main() {

    var appConfigs configs.AppConfig

    if appConfigs.TemplateCache == nil {
        fmt.Println("Creating template cache...")
        tc, err := renders.CreateTemplateCache()
        if err != nil {
            log.Fatal("Cannot create a cache for templates", err)
        }
        appConfigs.TemplateCache = tc
    }

    // Define handlers with tc
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        handlers.Home(w, r, appConfigs.TemplateCache["templates/home.page.html"])
    })

    http.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
        handlers.About(w, r, appConfigs.TemplateCache["templates/about.page.html"])
    })

    // Start the server and listen on port 8080
    fmt.Println(fmt.Sprintf("Starting server on port %s...", portNumber))
	err_listen := http.ListenAndServe(portNumber, nil)
	if err_listen != nil {
        fmt.Println("Error starting server:", err_listen)
    }
}