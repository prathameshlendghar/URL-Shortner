package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prathameshlendghar/URL-Shortner/api"
	"github.com/prathameshlendghar/URL-Shortner/internal/database"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error: Loading Environment variable failed")
	}

	database.ConnectDB()

	database.CreateTableIfNotExists()

	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "localhost"
	}
	hostport := os.Getenv("HOSTPORT")
	if hostport == "" {
		hostport = "8000"
	}

	router := api.RoutesSetup()

	fmt.Printf("Starting server at %v:%v\n", hostname, hostport)

	addressStr := fmt.Sprintf("%v:%v", hostname, hostport)
	log.Fatal(http.ListenAndServe(addressStr, router))
}
