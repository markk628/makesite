package main

import (
	"flag"
	"os"
	"html/template"
	"io/ioutil"
	"strings"
)

type post struct {
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

func renderTemplate(filename string, data string) {
	c := post{Content: data}
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	var err error
	err = t.Execute(os.Stdout , c)
	if err != nil {
		panic(err)
	}
}

func addExtHTML(filename string) string {
	ext := ".html"
	withExt := strings.Split(filename, ".")[0] + ext
	return withExt
}

func writeTemplateToFile(tmplName string, data string) {
	p := post{Content: readFile(data)}
	t := template.Must(template.New("template.tmpl").ParseFiles(tmplName))

	file := addExtHTML(data)
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, p)
	if err != nil {
		panic(err)
	}
}

func main() {
	filePtr := flag.String("file", "", "txt file to be converted to html file")
	flag.Parse()
	if *filePtr != "" {
		renderTemplate("template.tmpl", readFile(*filePtr))
		writeTemplateToFile("template.tmpl", *filePtr)
	} else {
		print("run what?")
	}
}