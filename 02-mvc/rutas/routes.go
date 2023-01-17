package rutas

import (
	"net/http"
	"paquetes/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Init)
	http.HandleFunc("/productos", controllers.Product)
}
