package models

import (
	"fmt"
	"log"
	"mycrud/config"

	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Id_usuario int16
	Nombre     string
	Correo     string
	Clave      string
}

func Hash_password(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14) //recomendado mas alta
	return string(bytes), err
}

func Check_password(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))

	if err != nil {
		panic("LAs contrasenas no coinciden")
	} else {
		return true
	}
}

func Cretae_user(name string, email string, password string) {
	encrypt_pass, _ := Hash_password(password)

	err, connect := config.Connect()

	if err != nil {
		panic(" Error al cifrar la contrasena ")
	}

	query, err := connect.Prepare("INSERT INTO usuario (nombre,correo,clave,activo)  values (?,?,?,1)")
	if err != nil {
		panic("Error al momento de crear usuario : ")
	} else {
		query.Exec(name, email, encrypt_pass)
		query.Close()
	}
}

func ReadUser(id string) Usuario {
	err, connect := config.Connect()

	if err != nil {
		panic(" Error al conectar la base de datos , ReadUsers")
	}

	row := connect.QueryRow("SELECT  id_usuario,nombre,correo,clave FROM usuario WHERE activo=1 AND id_usuario=?", id)

	connect.Close()

	var usuario Usuario
	row.Scan(&usuario.Id_usuario, &usuario.Nombre, &usuario.Correo, &usuario.Clave)
	fmt.Println("Usuario leido " + usuario.Nombre)
	return usuario
}

func ReadUsers() []Usuario {
	err, connect := config.Connect()

	if err != nil {
		panic(" Error al conectar la base de datos , ReadUsers")
	}

	rows, err := connect.Query("SELECT  id_usuario,nombre,correo FROM usuario WHERE activo=1")
	if err != nil {
		panic("Error al leer usuarios")
	}

	connect.Close()
	var Users []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.Id_usuario, &usuario.Nombre, &usuario.Correo)
		Users = append(Users, usuario)
	}
	return Users
}

func Delete_User(id string) {

	err, connect := config.Connect()

	if err != nil {
		panic(" Error al conectar con la base de datos ")
	}

	query, err := connect.Prepare("Update usuario set activo=? WHERE id_usuario= ?")
	if err != nil {
		panic("Error al momento de eliminar usuario : ")
	} else {
		query.Exec(0, id)
		query.Close()
	}
}

func Update_user(id, name, correo, pass string) {
	err, connect := config.Connect()

	if err != nil {
		panic(" Error al conectar con la base de datos ")
	}

	query, err := connect.Prepare("Update usuario set nombre=?,correo=?,clave=? WHERE id_usuario= ?")
	if err != nil {
		panic("Error al momento de actualizar usuario : ")
	}
	log.Println("Entra al else del update users")
	encrypt_pass, _ := Hash_password(pass)
	query.Exec(name, correo, encrypt_pass, id)
	query.Close()

}
