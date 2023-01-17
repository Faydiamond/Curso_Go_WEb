package routers

import (
	"templates/controllers"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Index)
	return router
}
