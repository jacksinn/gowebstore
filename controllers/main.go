package controllers

import (
	"net/http"
	"text/template"
	"os"
	"bufio"
	"strings"
	"github.com/jacksinn/gowebstore/viewmodels"
)

func Register(templates *template.Template) {
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request){
			requestedFile := req.URL.Path[1:] //remove first slash character
			template := templates.Lookup(requestedFile + ".html")

			var context interface{} = nil

			switch requestedFile {
			case "home":
				context = viewmodels.GetHome()
			case "categories":
				context = viewmodels.GetCategories()
			case "products":
				context = viewmodels.GetProducts()
			case "product":
				context = viewmodels.GetProduct()
			}



			if template != nil {
				template.Execute(w, context)
			} else {
				w.WriteHeader(404)
			}
		})

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)

}

func serveResource(w http.ResponseWriter, r *http.Request) {
	path := "public" + r.URL.Path

	var contentType string

	if strings.HasSuffix(path, ".css"){
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png"){
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}