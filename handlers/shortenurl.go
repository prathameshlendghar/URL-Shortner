package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/prathameshlendghar/URL-Shortner/internal/database"
	"github.com/prathameshlendghar/URL-Shortner/models"
)

func makeShortBase62(counter int64) string {
	base62 := "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
	var shortStr string = ""
	for counter > 0 {
		mod := counter % 62
		shortStr = shortStr + string(base62[mod])
		counter /= 62
	}
	return shortStr
}

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

	// parsedLongUrl := parsedUrl.String()

	// fmt.Println(parsedLongUrl)

	//TODO: Check for expiration is there or else give the expration

	//Take the counter from postgres database
	var counter int64 = database.GetCounter()

	//Create a short url in base62 format
	shortUniqueStr := makeShortBase62(counter)
	// fmt.Println(shortUniqueStr)

	//Store it inside the Database
	var dbStruct models.ShortUrlDB
	dbStruct.Id = counter
	dbStruct.LongUrl = requestBody.LongUrl
	dbStruct.ShortUrl = shortUniqueStr
	dbStruct.CreatedAt = time.Now()
	dbStruct.ExpireAt = dbStruct.CreatedAt.Add(48 * time.Hour)
	dbStruct.Tag = "Abc"

	database.InsertShortUrl(&dbStruct)

	// if requestBody.ExpireAt != "" {
	// 	dbStruct.ExpireAt = requestBody.ExpireAt
	// } else {
	// 	dbStruct.ExpireAt = time.Now() + time
	// }

	//Return the short url to the user in writer response

}
