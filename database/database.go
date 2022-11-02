package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var (
	DB_USERNAME string = "root"
	DB_PASSWORD string = ""
	DB_NAME     string = "go_miniproject"
)

func Connect() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@/%s?parseTime=true",
		DB_USERNAME,
		DB_PASSWORD,
		DB_NAME,
	)

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("ERROR SAAT MEMBUAT KONEKSI: %s", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("ERROR SAAT MENGHUBUNGKAN KE DATABASE: %s", err)
	}

	log.Println("Terhubung ke MySQL")

}
