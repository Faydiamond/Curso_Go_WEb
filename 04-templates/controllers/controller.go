package controllers

import (
	"html/template"
	"net/http"
)

/*
type User struct {
	Name    string
	Age     int16
	Active  bool
	IsAdmin bool
}*/

type Work struct {
	Name  string
	State bool
}

type Activities struct {
	User  string
	Works []Work
}

func Index(w http.ResponseWriter, r *http.Request) {
	render_view, err := template.ParseFiles("views/home.html", "views/layout.html", "views/navbar.html")

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	//user := User{Name: "Fabian", Age: 29, Active: true, IsAdmin: false}
	activities := Activities{"Fabian", []Work{
		{"Tender la cama ", true},
		{"Estudiar Ingles ", true},
		{"Lavar Loza Ingles ", false},
		{"Almorzar ", false},
	}}

	render_view.Execute(w, activities)
}
