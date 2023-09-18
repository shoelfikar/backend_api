package controller

import (
	"html/template"
	"net/http"

	"github.com/shoelfikar/golang_backend_api/helper"
)

type webbController struct{}

type WebController interface{
	LoginWebController(w http.ResponseWriter, r *http.Request)
	DashboardWebController(w http.ResponseWriter, r *http.Request)
}

func NewWebController() WebController {
	return &webbController{}
}

func (c *webbController) LoginWebController(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/login.html")

	helper.PanicIfError(err)

	t.ExecuteTemplate(w, "login.html", "Halaman Login")
}

func (c *webbController) DashboardWebController(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/dashboard.html")

	helper.PanicIfError(err)

	t.ExecuteTemplate(w, "dashboard.html", "Halaman Dashboard")
}