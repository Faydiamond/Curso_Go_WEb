package main

import (
	"fmt"
	"net/http"
	"paquetes/rutas"
)

func main() {
	rutas.Routes()
	fmt.Println("server in port 3000")
	http.ListenAndServe(":3000", nil)
}
