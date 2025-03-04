package handlers

import "net/http"

func DeleteURL(w http.ResponseWriter, r* http.Request){
	//Validate the type of the request DELETE
	//Validate the URL given is valid or not 
	//Also validate the mainURL passed is valid or not
	//slice the URL to just take the shorten base62 part
	//Search into DB in the indexed column
	//if present delete the main URL/other attribute and return some response
	//else return error
}