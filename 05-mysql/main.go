package main

import (
	"fmt"
	"mysql/db"
	"mysql/router"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Variables de entorno no cargadas")
	}
}

func main() {
	fmt.Println("server : 3000")
	connect := db.Connection()
	fmt.Println(connect)
	routes := router.Router()
	http.ListenAndServe(":3000", routes)
}
