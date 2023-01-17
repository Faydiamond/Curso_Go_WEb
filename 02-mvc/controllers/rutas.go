package controllers

import (
	"fmt"
	"net/http"
)

func Init(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Funcion for index page")
}

func Product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Query().Get("name"))
}
