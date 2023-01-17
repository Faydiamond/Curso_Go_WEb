package main

import (
	"fmt"
	"net/http"
	"templates/routers"
)

func main() {
	fmt.Println("Server is running in port : 3000")
	routes := routers.Routers()
	http.ListenAndServe(":3000", routes)
}
