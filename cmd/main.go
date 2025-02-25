package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prathameshlendghar/URL-Shortner/api"
)

func main() {
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "localhost"
	}
	hostport := os.Getenv("HOSTPORT")
	if hostname == "" {
		hostname = "8000"
	}

	router := api.RoutesSetup()

	fmt.Printf("Starting server at %v:%v\n", hostname, hostport)

	addressStr := fmt.Sprintf("%v:%v", hostname, hostport)
	log.Fatal(http.ListenAndServe(addressStr, router))
}
