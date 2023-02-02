package controllers

import (
	"go-scristi/pkg/utils"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type templateData struct {
	Title       string
	CurrentDate string
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(append(utils.GetListOfTemplates(), "html/pages/index.html")...)
	t.Execute(w, templateData{CurrentDate: strconv.Itoa(time.Now().Year()), Title: "Inicio - Sebastian Cristi"})
}
