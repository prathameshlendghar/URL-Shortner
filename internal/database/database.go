package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connectionString := fmt.Sprintf(os.Getenv("CONNECTION_STRING"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Println(connectionString)
	var err error
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatalf("Error: Unable to connect DB %v", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error: DB connection Failure")
	}
}

func CreateTableIfNotExists() {
	query := `CREATE TABLE IF NOT EXISTS url_data (id BIGSERIAL NOT NULL UNIQUE, shorturl VARCHAR(100) PRIMARY KEY, createdAt DATE, deleteAt DATE, Tag VARCHAR(30))`
	res, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating url_data DB table: %v", err)
	} else {
		fmt.Println(res)
	}
}

func GetCounter() {
	var nextVal int64
	query := "SELECT nextval('url_data_id_seq'::regclass)"
	err := DB.QueryRow(query).Scan(&nextVal)

	if err != nil {
		log.Fatalf("Error in fetching the next sequence id: %v", err)
	}
	fmt.Println(nextVal)
}
