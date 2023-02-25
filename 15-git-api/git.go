package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// TODO: Create a repo on Github using personal access token

type repoRequest struct {
	RepoName        string `json:"name"` // make sure field names matches to the json fields accepted by client
	RepoDescription string `json:"description"`
}

type repoResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Owner `json:"owner"`
}
type Owner struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Url   string `json:"url"`
}

type ErrorsGithub struct {
	Message          string           `json:"message"`
	DocumentationUrl string           `json:"documentation_Url"`
	Errors           []arrayErrGithub `json:"errors"`
}
type arrayErrGithub struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}

const token = `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`

func main() {
	repo := repoRequest{
		RepoName:        "test440",
		RepoDescription: "test-repo",
	}
	fmt.Println(create(repo))
}

func create(repo repoRequest) error {

	jsonData, err := json.Marshal(repo)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	//NewRequestWithContext construct the request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.github.com/user/repos", bytes.NewReader(jsonData))

	//setting the headers
	req.Header.Set("Content-Type", "application/json") // setting headers // we will send a json to the server

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		var gitErr ErrorsGithub
		err = json.Unmarshal(data, &gitErr)
		if err != nil {
			return err
		}
		fmt.Printf("failed \n %+v\n", gitErr)
		return errors.New("repo creation failed")
	}

	var result repoResponse
	err = json.Unmarshal(data, &result)

	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", result)
	return nil

}
