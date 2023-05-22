package main

import (
	"database/sql"
	"log"

	_ "github.com/jmrobles/h2go"
)

func main() {
	// db, err := sql.Open("h2", "h2://sa@localhost/Users/mhyoo/workspace/db/work/test?logging=trace")
	db, err := sql.Open("h2", "h2://sa@localhost/Users/mhyoo/workspace/db/work/test?logging=trace")
	if err != nil {
		log.Fatalf("Can't connet to H2 Database: %s", err)
	}

	// err = db.Ping()

	// if err != nil {
	// 	log.Fatalf("Can't ping to H2 Database: %s", err)
	// } else {
	// 	log.Printf("H2 Database connected")
	// }

	result, err := db.Exec("INSERT INTO test VALUES (?, ?)", 10, "test")

	// var name string
	// err = db.QueryRow("SELECT name FROM tutorials_tbl WHERE id = 1").Scan(&name)

	if err != nil {
		log.Fatalf("db sql err %s", err)
	}

	// log.Printf("name : %s", name)

	// sql.Result.RowsAffected()
	n, _ := result.RowsAffected()

	if n == 1 {
		log.Printf("1 row inserted.")
	}
	db.Close()
}
