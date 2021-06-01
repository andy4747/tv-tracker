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

//172.24.0.2
func Connect() *Queries {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbPassword,
		dbName,
		dbSSLMode,
	)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot ping the db: ", err)
	}
	//create all the schemas
	err = CreateTableIfNotExists(
		conn,
		usersTable,
		tokensTable,
	)

	if err != nil {
		log.Fatalln("cannot create table: ", err.Error())
	}

	db := New(conn)
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
