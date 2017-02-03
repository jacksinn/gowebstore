package main

import (
	"net/http"
	"text/template"
	"os"
	"github.com/jacksinn/gowebstore/controllers"
)

func main() {
	templates := populateTemplates()

	controllers.Register(templates)

	http.ListenAndServe(":8000", nil)
}



func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)

	defer templateFolder.Close()

	//all contents in one pass, no biggie, relatively small folder
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)

	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath + "/" + pathInfo.Name())
		}

	}

	//spread operator, get slice as input, handle otherwise
	result.ParseFiles(*templatePaths...)

	return result
}
