package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {

	db, _ := sql.Open("mysql", os.Getenv("DB_URL"))
	err := db.Ping()

	if err != nil {
		fmt.Println("Error en la conexion de la base de datos : ", err)
	} else {
		fmt.Println("Conexion Exitosa!")
	}
	return db
}
