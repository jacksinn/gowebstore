package controllers

import (
	"net/http"
	"github.com/jacksinn/gowebstore/viewmodels"
	"text/template"
	"github.com/gorilla/mux"
	"strconv"
)

type productController struct {
	template *template.Template
}

func (this *productController) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)

	if err == nil {
		vm := viewmodels.GetProduct(id)
		w.Header().Add("Content-Type", "text/html")
		this.template.Execute(w, vm)
	} else {
		w.WriteHeader(404)
	}
}
