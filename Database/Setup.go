package Database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Init(host, port, user, password, dbname string) *sql.DB {
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", ConnectionString)

	if err != nil {
		log.Fatalf("Database Cannot Connect, err: %v", err)
	}

	defer db.Close()

	err = db.Ping()

	if (err != nil) {
		fmt.Printf("Able to ping database!\n")
	} else {
		log.Fatalf("Cannot ping database! err: %v\n", err)
	}

	return db
}