package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
)

// PageVariables represents the data to be used in the template
type PageVariables struct {
	Title   string
	Message string
}

func generateDocument(t Template) string {
	// Get the template file
	templatePath := fmt.Sprintf("./templates/%s", t.Name)

	// Read the contents of the file into a string variable
	content, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	// Convert the byte slice to a string
	templateString := string(content)

	// Create the template engine
	tmpl, err := template.New(t.Name).Parse(templateString)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create a buffer that implements io.Writer
	var buffer bytes.Buffer

	// Execute the template, passing in the data
	err = tmpl.Execute(&buffer, t.Variables)
	if err != nil {
		log.Fatal(err.Error())
	}

	return buffer.String()
}
