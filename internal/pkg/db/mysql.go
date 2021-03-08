package database

import (
	"database/sql"
	"log"
)

var Db *sql.DB

// InitDB membuat koneksi dengan database
func InitDB() {
	// Use root:dbpass@tcp(172.17.0.2)/todolist, if you're using Windows.
	db, err := sql.Open("mysql", "syukur:syukurs4j4!!!@tcp(localhost)/todolist")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}
