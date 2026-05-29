package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prathameshlendghar/URL-Shortner/internal/auth"
	"github.com/prathameshlendghar/URL-Shortner/internal/url"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// This `godotenv.Load()` -> This function loads the environment variable into this process environment variables
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", slog.Any("error", err))
		os.Exit(1)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")

	// Setting up the relational Database empty connection pool -> This do not check even if the DB is reachable
	// DB connection works in the way that it will not have any conn and when some request needs DB access it will create one and maintain that connection in the connection pool for next set of
	// Open may just validate its arguments without creating a connection to the database. To verify that the data source name is valid, call [DB.Ping].
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		slog.Error("Error in validating database arguments", slog.Any("error", err))
		os.Exit(1)
	}
	defer dbConn.Close()

	// Checking if we have set up DB perfectly
	// -> opens first connection here to check if DB is reachable
	// Ping verifies a connection to the database is still alive, establishing a connection if necessary.

	err = dbConn.Ping()
	if err != nil {
		slog.Error("Error in establishing DB connection", slog.Any("error", err))
		os.Exit(1)
	}

	// This is the validator package that check the validity of Json payload in request by checking all required fields are passed or not by using reflection
	v := validator.New()

	// Setting up the web server -> Pointing towards the routes
	router := http.NewServeMux()
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Engine is running smoothly!"))
		if err != nil {
			log.Println("Error writing response: ", err)
		}
	})
	auth.SetupRoutes(router, dbConn, v) //Function where all routes are declared
	url.SetupRoutes(router, dbConn, v)  //Function where all routes are declared

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second, // Protects against slowloris attacks
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	slog.Info(fmt.Sprintf("Server starting on %s", srv.Addr))

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Error in starting server", slog.Any("error", err))
		os.Exit(1)
	}
}
