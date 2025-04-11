package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type application struct {
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := flag.String("dsn", os.Getenv("DB_STRING"), "MySQL data source name")

	flag.Parse()

	_, err = openDB(*dsn)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	app := application{}
	http.ListenAndServe(":4000", app.routes())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
