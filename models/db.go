package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbName     = os.Getenv("DB_NAME")
	dbSSLMode  = os.Getenv("DB_SSL_MODE")
)

func Connect() *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbPassword,
		dbName,
		dbSSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("cannot ping the db: ", err)
	}
	//create all the schemas
	err = CreateTableIfNotExists(
		db,
		usersTable,
		tokensTable,
		moviesTable,
		statusType,
	)

	if err != nil {
		log.Fatalln("cannot create table: ", err.Error())
	}
	return db
}

func CreateTableIfNotExists(db *sql.DB, tables ...string) error {
	for _, query := range tables {
		_, err := db.Exec(query)
		if err != nil {
			log.Println("cannot create table: ", err.Error())
			return err
		}
	}
	return nil
}
