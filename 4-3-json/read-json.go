package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type user struct {
	Name  string          `json:"first_name"` // make sure name in field tag is same as the json
	Roles map[string]bool `json:"perms"`
}

func main() {
	jsonData, err := os.ReadFile("test.json")
	if err != nil {
		log.Fatalln(err)
	}

	//var p *int
	//var p1 **int = &p // not a good idea

	var u []user
	err = json.Unmarshal(jsonData, &u) // Unmarshal is one of the exception
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v", u)

}
