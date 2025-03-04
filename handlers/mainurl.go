package handlers

import "net/http"

func GetMainURL(w http.ResponseWriter, r *http.Request) {
	//Validate the type of the request
	//Validate the URL given is valid or not 
	//slice the URL to just take the shorten base62 part
	//Search into DB in the indexed column
	//if present return the main URL+ other attribute
	//else return error 
}
