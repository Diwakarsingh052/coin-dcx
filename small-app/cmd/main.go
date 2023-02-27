package main

import (
	"net/http"
	"small-app/handlers"
)

// /data?id=2

func main() {

	http.HandleFunc("/user", handlers.GetUser)
	http.ListenAndServe(":8080", nil)

}
