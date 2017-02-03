package controllers

import (
	"net/http"
	"github.com/jacksinn/gowebstore/viewmodels"
	"text/template"
	"github.com/gorilla/mux"
	"strconv"
)

type categoriesController struct {
	template *template.Template
}

func (cc *categoriesController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.GetCategories()

	w.Header().Add("Content-Type", "text/html")

	cc.template.Execute(w, vm)
}

type categoryController struct {
	template *template.Template
}

func (this *categoryController) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)
	if err == nil {
		vm := viewmodels.GetProducts(id)

		w.Header().Add("Content-Type", "text/html")

		this.template.Execute(w, vm)
	} else {
		w.WriteHeader(404)
	}

}