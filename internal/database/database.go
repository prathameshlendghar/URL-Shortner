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

func InsertShortUrl(urlDetails *models.ShortUrlDB) (models.ShortUrlResp, error) {
	var resp models.ShortUrlResp
	query := `INSERT INTO url_data(id, short_code, original_url, createdAt, deleteAt, tag) 
				values ($1, $2, $3, $4, $5, $6) RETURNING id, short_code, original_url, createdAt, deleteAt, tag`
	var id int
	err := DB.QueryRow(
		query,
		urlDetails.Id,
		urlDetails.ShortUrl,
		urlDetails.LongUrl,
		urlDetails.CreatedAt.Format("2006-01-02"),
		urlDetails.ExpireAt.Format("2006-01-02"),
		urlDetails.Tag).Scan(&id, &resp.ShortUrl, &resp.LongUrl, &resp.CreatedAt, &resp.ExpiresAt, &resp.Tag)

	return resp, err
}

func FetchLongUrl(shortUrl string) (string, error) {
	var resp string
	query := `SELECT original_url FROM url_data WHERE short_code = $1`
	err := DB.QueryRow(query, shortUrl).Scan(&resp)

	return resp, err
}

//Keeping the above and this funcion different because of latency and network load
//Like why to send entire data when only redirection is required

func FetchUrlInfo(shortUrl string) (models.ShortUrlDB, error) {
	var resp models.ShortUrlDB
	query := `SELECT original_url, short_code, createdat, deleteat, tag FROM url_data WHERE short_code = $1`
	err := DB.QueryRow(query, shortUrl).Scan(&resp.LongUrl, &resp.ShortUrl, &resp.CreatedAt, &resp.ExpireAt, &resp.Tag)

	return resp, err
}
