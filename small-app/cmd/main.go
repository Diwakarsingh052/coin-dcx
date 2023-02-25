package main

import (
	"net/http"
	"small-app/handlers"
)

// /user?id=2

func main() {

	http.HandleFunc("/user", handlers.GetUser)
	http.ListenAndServe(":8080", nil)

}
