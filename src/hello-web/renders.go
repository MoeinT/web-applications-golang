package main

import (
    "fmt"
    "net/http"
    "html/template"
	"path/filepath"
	"log"
)


func RenderTemplate(w http.ResponseWriter, data interface{}, templateFile string) {
	
	// Get the cached templates
	cachedTmpl, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Get the cached template
	parsedTemplates, InMap:= cachedTmpl[templateFile]
	if !InMap {
		log.Fatal("cached template not found")
	}

	// Execute the cached template
	err_exec := parsedTemplates.Execute(w, data)
	if err_exec != nil {
		log.Fatal(err_exec)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	var cachedTmpl = make(map[string]*template.Template)

	files, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		fmt.Println("Error on files Glob:", err)
		return cachedTmpl, err
	}

	for _, file := range files {
		// Parse the file
		fileParsed, err := template.ParseFiles(file)
		if err != nil {
			fmt.Println("Error parsing files template:", err)
			return cachedTmpl, err
		}

		// Get layout files
		layouts, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			fmt.Println("Error on layout glob:", err)
			return cachedTmpl, err
		}

		// Add layout files to the parsed file
		if len(layouts) > 0 {
			fileParsed, err = fileParsed.ParseFiles(layouts...)
			if err != nil {
				fmt.Println("Error parsing layout files:", err)
				return cachedTmpl, err
			}
		}

		cachedTmpl[file] = fileParsed
	}

	return cachedTmpl, nil
}