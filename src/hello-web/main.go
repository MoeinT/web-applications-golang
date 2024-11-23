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
        appConfigs.UseCache = false
    }

    // Pass the application configs to the renders package
    renders.NewTemplateCache(appConfigs)
    // Pass the application configs to the handlers package
    handler := &handlers.App{AppConfig: &appConfigs}

    // Call the Home and About handlers from the handlers package
    http.HandleFunc("/", handler.Home)
    http.HandleFunc("/about/", handler.About)

    // Start the server and listen on port 8080
    fmt.Println(fmt.Sprintf("Starting server on port %s...", portNumber))
	err_listen := http.ListenAndServe(portNumber, nil)
	if err_listen != nil {
        fmt.Println("Error starting server:", err_listen)
    }
}