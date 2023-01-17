package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " Pagina de inicio! ")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " Acerca de nosotros! ")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " Contactenos ")
}

func Pokemons(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " Listado de pokemones! ")
}

func Pokemon(w http.ResponseWriter, r *http.Request) {
	id_pokemon := r.URL.Query().Get("id")
	pokemon := r.URL.Query().Get("name")
	fmt.Fprintln(w, id_pokemon+" "+pokemon)
}
