package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func conectar_postgres() {
	connStr := "user=tu_usuario password=tu_contraseña dbname=tu_bd host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a PostgreSQL:", err)
	}

	fmt.Println("Conexión exitosa a PostgreSQL")
}