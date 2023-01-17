package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Init(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Inicio")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Acerca de")
}

func Obtener_libro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name_book := params["name_book"]
	fmt.Fprintln(w, name_book)
}
