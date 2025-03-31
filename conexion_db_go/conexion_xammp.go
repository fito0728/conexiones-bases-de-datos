package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func conectar_xampp() {
	connStr := "usuario:contraseña@tcp(127.0.0.1:3306)/tu_bd"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a MySQL:", err)
	}

	fmt.Println("Conexión exitosa a MySQL en XAMPP")
}