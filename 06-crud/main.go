package main

import (
	"fmt"
	"log"
	"mycrud/routes"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("No se pudieron cargar las variables de entorno")
	} else {
		fmt.Println("Se cargaron correctamente las variables de entorno")
	}
}

func main() {
	fmt.Println("Serve in port 3000")
	route := routes.Routes()

	http.ListenAndServe(":3000", route)
}
