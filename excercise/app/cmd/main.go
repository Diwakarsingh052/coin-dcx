package main

import (
	"net/http"
)

// /data?id=2

func main() {

	http.HandleFunc("/user")
	http.ListenAndServe(":8080", nil)

}
