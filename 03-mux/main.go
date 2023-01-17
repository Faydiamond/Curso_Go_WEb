package main

import (
	"fmt"
	"muxx/routers"
	"net/http"
)

func main() {
	fmt.Println("server running in port 3000")
	routes := routers.Routes()
	http.ListenAndServe(":3000", routes) // se llaman a los mux
}
