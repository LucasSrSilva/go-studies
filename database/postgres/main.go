package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connStr := "postgres://admin:admin@localhost:5432/gopostgres"
	var postgresVer string
	var resultado string
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = db.QueryRow("Select version()").Scan((&postgresVer))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows, err := db.Query("SELECT name FROM cursos")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {

		err := rows.Scan(&resultado)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Lista:  %s\n", resultado)
	}
	fmt.Println("Conectado")
	fmt.Printf("Versão do db: %s\n", postgresVer)

}
