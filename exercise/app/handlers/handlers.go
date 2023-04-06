package handlers

import (
	"app/data/user"
	"encoding/json"
	"log"
	"net/http"
)

// /user?user_id=2

// GetUser is entry point for /user endpoint
// think how would you handle the request when someone hit this endpoint
func GetUser(w http.ResponseWriter, r *http.Request) {
	// this line set your  ContentType as json
	w.Header().Set("Content-Type", "application/json")

	//fetching the variable from query
	userIdString := r.URL.Query().Get("user_id")

	//converting it to make sure it is a valid uint64 // strconv.parseuint

	if err != nil {
		log.Println(err)
		appErr := map[string]string{"Message": http.StatusText(http.StatusBadRequest)}

		w.WriteHeader(http.StatusBadRequest) // setting error status code
		json.NewEncoder(w).Encode(appErr)    // converting map to json and sending back to the client using responseWritet
		return

		////signal with text based error
		//http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		//return // don't forget the return
	}

	//fetching the data with the userId provided
	u, err := user.FetchUser(userId)
	if err != nil {

		//handle error like previous case

		//http.Error(w, err.Error(), http.StatusBadRequest)
		//return
	}

	//encode the user
	json.NewEncoder(w).Encode()
	//b, _ := json.Marshal(u)
	//w.Write(b)
}
