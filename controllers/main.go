package controllers

import (
	"net/http"
	"text/template"
	"os"
	"bufio"
	"strings"
	//"github.com/jacksinn/gowebstore/viewmodels"
	"github.com/gorilla/mux"
)

func Register(templates *template.Template) {
	router := mux.NewRouter()

	//Home Page
	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	router.HandleFunc("/home", hc.get)

	//Categories Page
	cc := new(categoriesController)
	cc.template = templates.Lookup("categories.html")
	router.HandleFunc("/categories", cc.get)

	categoryController := new(categoryController)
	categoryController.template = templates.Lookup("products.html")
	router.HandleFunc("/categories/{id}", categoryController.get)

	productController := new(productController)
	productController.template = templates.Lookup("product.html")
	router.HandleFunc("/product/{id}", productController.get)

	http.Handle("/", router)
	//Handling Images and CSS
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