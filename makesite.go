package main

import (
	"fmt"
	"io/ioutil"
	"html/template"
	"os"
)

type post struct {
	User string
	Content string
}

func readFile(templateName string) string {
	fileContents, err := ioutil.ReadFile(templateName)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
        // value that we don’t know how to (or want to) handle. This example
        // panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	return string(fileContents)
}

func renderTemplate(content string) *template.Template {
	path:= []string {
		"template.tmpl",
	}

	templateFile := template.Must(template.New("template.tmpl").ParseFiles(path...))
	err := templateFile.Execute(os.Stdout, post{User: "John Doe", Content: content})
	if err != nil {
		panic(err)
	}
	return templateFile
}

func main() {
	fmt.Println("Evenin lads")

	content := readFile("first-post.txt")
	t := renderTemplate(content)
	print(t)
}