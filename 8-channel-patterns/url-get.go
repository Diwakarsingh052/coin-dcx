package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var url = []string{
	`https://pkg.go.dev/`,
	`https://github.com/`,
	`abc.com/1234`,
}

type response struct {
	url  string
	resp *http.Response
	err  error
}

var wg = &sync.WaitGroup{}

func main() {
	doGetRequest(url)
	wg.Wait()
}

func doGetRequest(urls []string) {

	respChan := make(chan response, len(urls)) // buffered channel
	wgGet := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range urls {

			// keeping track of get requests that are going to start in the next line
			//wgGet.Add(1)

			//fanning out go routines // one task = one goroutine
			wgGet.Add(1)
			go func(url string) {

				defer wgGet.Done()

				resp, err := http.Get(url)

				r := response{
					url:  url,
					resp: resp,
					err:  err,
				}

				respChan <- r //sending the resp struct to respCh

			}(v)

		}
		//waiting for go routines to finish the get request task
		wgGet.Wait()
		close(respChan)
		// when channel is closed no more send can happen // only recv is possible

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//range would stop when channel is closed
		for r := range respChan { // recv over the channel until senders are sending values or channel is not closed

			if r.err != nil {
				log.Println(r.err)
				continue
			}

			bytes, err := io.ReadAll(r.resp.Body)
			if err != nil {
				log.Println(err)
				continue
			}
			defer r.resp.Body.Close()

			if r.resp.StatusCode > 299 {
				log.Printf("Response failed with status code: %d and\nbody: %s\n", r.resp.StatusCode, bytes)
				continue
			}

			fmt.Println(r.url, r.resp.Status)
		}
	}()

}
