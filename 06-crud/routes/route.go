package routes

import (
	"mycrud/handlers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	userRouter := router.PathPrefix("/users").Subrouter()
	//router.HandleFunc("/", handlers.Index)
	userRouter.HandleFunc("/", handlers.ReadUsers).Methods("GET")
	userRouter.HandleFunc("/create", handlers.Create_user).Methods("POST", "GET")
	userRouter.HandleFunc("/delete", handlers.Delete_user).Methods("GET")
	userRouter.HandleFunc("/update", handlers.Update_user).Methods("GET", "POST")
	return router
}
