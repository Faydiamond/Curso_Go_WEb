package main

import (
	"fmt"
	"net/http"
	"pokes/handlers"
)

func main() {
	fmt.Println("port in use 3000!")
	handlers.Handlers()
	http.ListenAndServe(":3000", nil)
}
