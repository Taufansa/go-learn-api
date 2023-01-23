package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetPostgresDB() (*sql.DB, error) {
	host := "localhost"    //os.Getenv("POSTGRES_HOST")
	user := "postgres"     //os.Getenv("POSTGRES_USER")
	password := "postgres" //os.Getenv("POSTGRES_PASSWORD")
	dbName := "go-learn"   //os.Getenv("POSTGRES_DB")

	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

	db, err := createConnection(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createConnection(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
