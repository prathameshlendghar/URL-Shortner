package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/prathameshlendghar/URL-Shortner/models"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		response := map[string]string{"error": "Only POST Method is Allowed"}
		json.NewEncoder(w).Encode(response)

		return
	}

	//Check if the request body is valid or not
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var requestBody models.NewUrlReq
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "error: Unable to parse request body", http.StatusBadRequest)
		return
	}
	//Check the validity of long url
	longUrl := requestBody.LongUrl
	parsedUrl, err := url.Parse(longUrl)

	if err != nil {
		errorstr := fmt.Sprintf("error: %v", err)
		http.Error(w, errorstr, http.StatusBadRequest)
		return
	}

	if parsedUrl.Scheme != "" && parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"error": "Supported methods are http and https"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "http"
	}

	parsedUrl.RawQuery = url.QueryEscape(parsedUrl.RawQuery)

	parsedLongUrl := parsedUrl.String()

	fmt.Println(parsedLongUrl)



	//TODO: Check for expiration is there or else give the expration

	//Take the counter from postgres database

	//Create a short url in base62 format

	//Store it inside the Database

	//Return the short url to the user in writer response

}
