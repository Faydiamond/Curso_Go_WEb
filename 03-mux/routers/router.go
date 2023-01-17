package routers

import (
	"muxx/controllers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Init)
	router.HandleFunc("/about", controllers.About).Methods("GET")
	router.HandleFunc("/libro/{name_book}", controllers.Obtener_libro).Methods("GET")
	return router
}
