package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (error, *sql.DB) {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return err, nil
	} else {
		fmt.Println("Felicidades conexion a la base de datos ")
		return nil, db
	}
}
