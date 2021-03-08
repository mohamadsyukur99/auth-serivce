package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	setting "github.com/mohamadsyukur99/auth-service/internal/pkg/setting"
)

var Db *sql.DB

// InitDB membuat koneksi dengan database
func InitDB() {
	setting, _ := setting.LoadConfiguration("setting.json")
	database := setting.Db
	// Use root:dbpass@tcp(172.17.0.2)/todolist, if you're using Windows.
	// db, err := sql.Open("mysql", "syukur:syukurs4j4!!!@tcp(localhost)/todolist")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", database.User, database.Password, database.Host, database.Database))

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}
