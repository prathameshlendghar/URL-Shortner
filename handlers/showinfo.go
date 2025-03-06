package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/prathameshlendghar/URL-Shortner/internal/database"
	"github.com/prathameshlendghar/URL-Shortner/models"
	"github.com/prathameshlendghar/URL-Shortner/utils"
)

func ShowInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSONUtils(w, http.StatusMethodNotAllowed, "Error: Only Get Requests are Allowed")
		return
	}

	var requestData models.GetInfoReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)

	if err != nil {
		errStr := fmt.Sprintf("Error: Unable to Decode Request Body: %v", err)
		utils.WriteJSONUtils(w, http.StatusBadRequest, errStr)
		return
	}

	parsedRequestUrl, err := url.Parse("http://" + requestData.ShortUrl)
	if err != nil {
		utils.WriteJSONUtils(w, http.StatusBadRequest, "Error: Unable to Decode Request Body1")
		return
	}

	shortUrlCode := parsedRequestUrl.Path
	shortUrlCode = shortUrlCode[1:]
	fmt.Println(shortUrlCode)
	resp, err := database.FetchUrlInfo(shortUrlCode)
	// fmt.Println(resp)
	if err != nil {
		str := fmt.Sprintf("Error: Unable to fetch Info: %v", err)
		utils.WriteJSONUtils(w, http.StatusInternalServerError, str)
		return
	}

	utils.WriteJSONUtils(w, http.StatusOK, resp)
}
