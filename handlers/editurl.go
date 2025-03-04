package handlers

import "net/http"

func EditURL(w http.ResponseWriter, r* http.Request){
	//Validate the type of the request PUT
	//Validate the URL given is valid or not 
	//Also validate the mainURL passed is valid or not
	//slice the URL to just take the shorten base62 part
	//Search into DB in the indexed column
	//if present edit the main URL/other attribute and return the edited info
	//else return error
}