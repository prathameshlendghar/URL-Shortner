package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/prathameshlendghar/URL-Shortner/models"
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
	query := `CREATE TABLE IF NOT EXISTS url_data (id BIGSERIAL NOT NULL UNIQUE, short_code VARCHAR(10) PRIMARY KEY, original_url TEXT, createdAt DATE, deleteAt DATE, tag VARCHAR(30))`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating url_data DB table: %v", err)
	}
}

func GetCounter() int64 {
	var nextVal int64
	query := "SELECT nextval('url_data_id_seq'::regclass)"
	err := DB.QueryRow(query).Scan(&nextVal)

	if err != nil {
		log.Fatalf("Error in fetching the next sequence id: %v", err)
	}
	return nextVal
}

func InsertShortUrl(urlDetails *models.ShortUrlDB) {
	query := `INSERT INTO url_data(id, short_code, original_url, createdAt, deleteAt, tag) values ($1, $2, $3, $4, $5, $6)`
	_, err := DB.Query(query, urlDetails.Id, urlDetails.ShortUrl, urlDetails.LongUrl, urlDetails.CreatedAt.Format("2006-01-02"), urlDetails.ExpireAt.Format("2006-01-02"), urlDetails.Tag)
	if err != nil {
		log.Fatal(err)
	}
}
