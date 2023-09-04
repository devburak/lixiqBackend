package utils

import (
	"bytes"
	"html/template"
	"log"
	"path/filepath"
)

// LoadTemplate loads an HTML template from the specified file path and parses it with data.
func LoadTemplate(tmplPath string, data interface{}) (string, error) {
	basePath := filepath.Dir("internal/template/")
	fullPath := filepath.Join(basePath, tmplPath)

	tmpl, err := template.ParseFiles(fullPath)
	if err != nil {
		log.Println("Path:", fullPath)
		log.Fatal("Error loading template :", err)
		return "", err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		log.Fatal("Error executing template with data :", err)
		return "", err
	}

	return tpl.String(), nil
}
