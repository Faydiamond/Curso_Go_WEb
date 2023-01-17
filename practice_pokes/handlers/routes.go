package handlers

import (
	"net/http"
	"pokes/controllers"
)

func Handlers() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/about", controllers.About)
	http.HandleFunc("/contact", controllers.Contact)
	http.HandleFunc("/pokemons", controllers.Pokemons)
	http.HandleFunc("/pokemon", controllers.Pokemon)
}
