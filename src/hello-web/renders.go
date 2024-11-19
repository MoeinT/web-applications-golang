package main

import (
    "fmt"
    "net/http"
    "html/template"
)

var cachedTmpl = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, data interface{}, templateFile string) {
	
	_, inMap := cachedTmpl[templateFile]

	if !inMap {
		CreateTemplateCache(templateFile)
	}

	if err := cachedTmpl[templateFile].Execute(w, data); err != nil {
		fmt.Println("Error executing template:", err)
	}
}

func CreateTemplateCache(templateFile string) error {
	templates := []string{
		templateFile,
		"templates/base.layout.html",
	}

	tmplParsed, errorParsed := template.ParseFiles(templates...)
	if errorParsed != nil {
		return errorParsed
	}

	cachedTmpl[templateFile] = tmplParsed
	fmt.Println("Template added to cache:", templateFile)

	return nil
}