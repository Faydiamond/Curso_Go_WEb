package router

import (
	"fmt"
	"mysql/db"
	"net/http"

	"github.com/gorilla/mux"
)

type Producto struct {
	id         int16
	nombre     string
	existencia int16
	activo     int16
}

func insert() {
	//db := db.Connection()
	query := " INSERT INTO productos (nombre, existencia,activo) VALUES (?,?,?) "

	/*
		vars := mux.Vars(r)

		nombre := vars["nombre"]
		existencia := vars["existencia"]
		activo := 1
		res, err := db.Exec(query, nombre, existencia, activo)

		if err != nil {
			fmt.Println("Registro no se inserto")
		} else {
			fmt.Println("Insertado:  ", res)
		}*/
	fmt.Println(query)
}

func Router() *mux.Router {
	db := db.Connection()
	router := mux.NewRouter()
	router.HandleFunc("/productos", func(w http.ResponseWriter, r *http.Request) {
		query := " SELECT  id_producto,nombre, existencia,activo FROM productos "
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println(" Registro no encontrado ")
		} else {
			fmt.Fprintln(w, "<ul>")

			var productos []Producto
			for rows.Next() {
				var producto Producto
				rows.Scan(&producto.id, &producto.nombre, &producto.existencia, &producto.activo)
				productos = append(productos, producto)
				fmt.Fprintln(w, productos)

			}

			fmt.Fprintln(w, "</ul>")
		}
	})
	return router
}
