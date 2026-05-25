package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/prathameshlendghar/URL-Shortner/internal/database"
	"github.com/prathameshlendghar/URL-Shortner/models"
	"github.com/prathameshlendghar/URL-Shortner/utils"
)

func EditURL(w http.ResponseWriter, r *http.Request) {
	//Validate the type of the request PUT
	if r.Method != http.MethodPut {
		utils.WriteJSONUtils(w, http.StatusMethodNotAllowed, "Error: Only Put Requests are Allowed")
		return
	}

	var requestData models.UpdateReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)

	if err != nil {
		errStr := fmt.Sprintf("Error: Unable to Decode Request Body: %v", err)
		utils.WriteJSONUtils(w, http.StatusBadRequest, errStr)
		return
	}

	//Validate the URL given is valid or not

	//Also validate the mainURL passed is valid or not

	var updateReqDB models.UpdateReqDB

	if requestData.LongUrl != "" {
		longUrl, err := utils.ValidateLongUrl(requestData.LongUrl)
		if err != nil {
			errStr := fmt.Sprintf("Error: Invalid Long Url Type: %v", err)
			utils.WriteJSONUtils(w, http.StatusBadRequest, errStr)
		}
		updateReqDB.LongUrl = &longUrl
	}

	if requestData.ExpireAfter != 0 {
		expireAt := time.Now().Add(24 * time.Duration(requestData.ExpireAfter) * time.Hour)
		updateReqDB.ExpireAt = &expireAt
	}
	updateReqDB.Tag = &requestData.Tag

	//slice the URL to just take the shorten base62 part
	shortUrl, err := utils.ValidateShortUrl(requestData.ShortUrl)
	if err != nil {
		errStr := fmt.Sprintf("Error: Unable to Decode Short Url: %v", err)
		utils.WriteJSONUtils(w, http.StatusBadRequest, errStr)
		return
	}

	//Search into DB in the indexed column
	resp, err := database.UpdateUrlInfo(updateReqDB, shortUrl)
	if err != nil {
		utils.WriteJSONUtils(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSONUtils(w, http.StatusAccepted, resp)
	//if present edit the main URL/other attribute and return the edited info
	//else return error
}
