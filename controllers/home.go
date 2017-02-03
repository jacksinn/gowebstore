package controllers

import (
	"net/http"
	"github.com/jacksinn/gowebstore/viewmodels"
	"text/template"
)

type homeController struct {
	template *template.Template
}

func (hc *homeController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.GetHome()

	w.Header().Add("Content-Type", "text/html")

	hc.template.Execute(w, vm)
}