package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	msg := "metodo invalido"
	if method == "GET" {
		msg = "metodo valido"
	}

	fmt.Fprintln(w, msg)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1 style='color:green'> Contact </h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1 style='color:black'> About </h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/BOUT", about)
	fmt.Println("App server port :3000")
	http.ListenAndServe(":3000", nil)

}
