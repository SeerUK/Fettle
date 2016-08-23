package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("==> Connecting to database...")

	db, err := sql.Open("mysql", "root:fettle@tcp(172.16.203.145:3306)/fettle")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("    Connection successful.")
}
