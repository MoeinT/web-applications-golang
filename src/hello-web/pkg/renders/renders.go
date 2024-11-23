package renders

import (
    "fmt"
    "net/http"
    "html/template"
	"path/filepath"
	"log"
	"bytes"
)


func RenderTemplate(w http.ResponseWriter, data interface{}, parsedTemplates *template.Template) {

	buff := new(bytes.Buffer)

	err_buff := parsedTemplates.Execute(buff, data)
	if err_buff != nil {
		log.Fatal(err_buff)
	}

	_, err_write_buff := buff.WriteTo(w)
	if err_write_buff != nil {
		log.Fatal(err_write_buff)
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