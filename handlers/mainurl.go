package handlers

import (
	"net/http"
	"strings"

	"github.com/prathameshlendghar/URL-Shortner/internal/database"
	"github.com/prathameshlendghar/URL-Shortner/utils"
)

func GetMainURL(w http.ResponseWriter, r *http.Request) {
	//Validate the type of the request
	if r.Method != http.MethodGet {
		utils.WriteJSONUtils(w, http.StatusMethodNotAllowed, "Error: Only Get Request is allowed")
		return
	}
	//Validate the URL given is valid or not
	shortUrlPath := r.URL.Path
	//slice the URL to just take the shorten base62 part
	shortUrlPath = shortUrlPath[1:]

	//Search into DB in the indexed column
	resp, err := database.FetchLongUrl(shortUrlPath)
	if err != nil {
		utils.WriteJSONUtils(w, http.StatusInternalServerError, "Error: Unable to fetch the shortUrl")
		return
	}

	if !strings.HasPrefix(resp, "http://") && !strings.HasPrefix(resp, "https://") {
		resp = "http://" + resp
	}
	http.Redirect(w, r, resp, http.StatusMovedPermanently)

}
