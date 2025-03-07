package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prathameshlendghar/URL-Shortner/internal/database"
	"github.com/prathameshlendghar/URL-Shortner/models"
	"github.com/prathameshlendghar/URL-Shortner/utils"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		response := map[string]string{"error": "Only POST Method is Allowed"}
		utils.WriteJSONUtils(w, http.StatusMethodNotAllowed, response)
		return
	}

	//Check if the request body is valid or not
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var requestBody models.NewUrlReq
	if err := decoder.Decode(&requestBody); err != nil {
		utils.WriteJSONUtils(w, http.StatusBadRequest, "error: Unable to parse request body")
		return
	}

	//Check the validity of long url
	longUrl := requestBody.LongUrl
	parsedLongUrl, err := utils.ValidateLongUrl(longUrl)
	if err != nil {
		utils.WriteJSONUtils(w, http.StatusBadRequest, err)
	}

	//Check for expiration is there or else give the expration
	if requestBody.ExpireAfter == 0 {
		exp_period, err := strconv.Atoi(os.Getenv("DEFAULT_EXPIRATION_PERIOD"))
		if err != nil {
			errStr := "Unable to read expireAt's duration"
			http.Error(w, errStr, http.StatusBadRequest)
		}
		requestBody.ExpireAfter = int32(exp_period)
	}

	//Take the counter from postgres database
	var counter int64 = database.GetCounter()

	//Create a short url in base62 format
	shortUniqueStr := utils.MakeShortBase62(counter)

	//Store it inside the Database and return response

	dbStruct := models.ShortUrlDB{
		Id:        counter,
		LongUrl:   parsedLongUrl,
		ShortUrl:  shortUniqueStr,
		CreatedAt: time.Now(),
		ExpireAt:  time.Now().Add(24 * time.Duration(requestBody.ExpireAfter) * time.Hour),
		Tag:       requestBody.Tag,
	}

	resp, _ := database.InsertShortUrl(&dbStruct)
	resp.ShortUrl = os.Getenv("SHORTURL_HOST") + "/" + resp.ShortUrl

	utils.WriteJSONUtils(w, http.StatusAccepted, resp)
}
