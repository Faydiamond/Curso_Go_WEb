package handlers

import (
	"fmt"
	"html/template"
	"log"
	"mycrud/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, " Inicio ")
}

func Create_user(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method)
	//hacer un if
	if r.Method == "POST" {
		log.Println("Entra al metodo post para crear usuarios")
		name := r.FormValue("name")
		email := r.FormValue("email")
		pass := r.FormValue("pass")

		models.Cretae_user(name, email, pass)
		http.Redirect(w, r, "/users/", http.StatusFound) //Redireccionar
	}

	view, err := template.ParseFiles("views/create.html", "views/layaout.html")

	//http.Redirect(w, r, "/users/", http.StatusFound) //Redireccionar
	if err != nil {
		log.Fatal("No encuentro la vista creacion de usuario")
	} else {
		view.Execute(w, nil)
	}
}

func ReadUsers(w http.ResponseWriter, r *http.Request) {

	log.Println("ReadUsers: ", r.Method)
	//hacer un if

	//models.Cretae_user("Clara", "Clara24@gmail.com", "/*845111212121212")
	view, err := template.ParseFiles("views/read.html", "views/layaout.html")
	if err != nil {
		log.Fatal("No encuentro la vista Read usuario")
	} else {
		usuarios := models.ReadUsers()
		log.Println("usuarios :", usuarios)
		view.Execute(w, usuarios)
	}
}

func Delete_user(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id_usuario")
	models.Delete_User(id)
	http.Redirect(w, r, "/users/", http.StatusFound) //Redireccionar
}

func Update_user(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id_usuario")
	fmt.Println("ESte es el id por actualizar " + id)

	if r.Method == http.MethodPost {
		log.Println("Entra a actualizar usuario")
		nombre := r.FormValue("name")
		correo := r.FormValue("email")
		clave := r.FormValue("pass")

		models.Update_user(id, nombre, correo, clave)
		http.Redirect(w, r, "/users/", http.StatusFound) //Redireccionar

	}
	view, err := template.ParseFiles("views/updateUser.html", "views/layaout.html")

	//http.Redirect(w, r, "/users/", http.StatusFound) //Redireccionar
	if err != nil {
		log.Fatal("No encuentro la vista creacion de usuario")
	}
	fmt.Println("Metodo GET update users")
	usuario := models.ReadUser(id)
	log.Println("Que mierda ve  ", usuario)
	view.Execute(w, usuario)

}
