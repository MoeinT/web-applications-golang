package main

import (
    "fmt"
    "net/http"
    "html/template"
)

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