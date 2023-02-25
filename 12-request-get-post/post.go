package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	user := map[string]string{"first_name": "raj"}
	u, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//https://api.github.com/user/repos
	//NewRequestWithContext construct the request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://httpbin.org/post", bytes.NewReader(u))

	//setting the headers
	req.Header.Set("Content-Type", "application/json") // setting headers // we will send a json to the server

	//doing the request to the remote server
	resp, err := http.DefaultClient.Do(req)
	//err is going to tell us if we are not able to call the remote service
	// it will not indicate any errors that might have happened while exec the req
	if err != nil {
		log.Fatalln(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode > 299 {
		log.Fatalln(string(b))
	}
	log.Println(string(b))

}
