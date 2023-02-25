package main

import (
	"fmt"
	"net/http"
)

func main() {
	//registering the route // and registering the handler which will handle the req
	http.HandleFunc("/home", HomePage)

	// this will start the server //
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}

// HomePage
// http.ResponseWriter is used to write response back to the client ,
// http.Request is used to fetch any request specific details like json, body , or anything related to request data
func HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "this is my home page")
	//w.Write([]byte("this is my home page"))
	fmt.Println("this is my home page")
	fmt.Println(r.URL)
}
