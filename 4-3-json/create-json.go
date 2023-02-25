package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Perms map[string]bool
type person struct {
	FirstName string `json:"first_name"` //field tags // set the field name to the name specified for the json
	Password  string `json:"-"`          //- means ignore this field in json output
	Perms     `json:"perms,omitempty"`   //omitempty // if the field is null then avoid in json output
}

func main() {
	users := []person{
		{
			FirstName: "Roy",
			Password:  "abc",
			Perms:     Perms{"admin": true},
		},
		{
			FirstName: "John",
			Password:  "qwe",
		},
		{
			FirstName: "Bruce",
			Password:  "efg",
			Perms:     Perms{"write": false},
		},
	}

	//jsonData, err := json.Marshal(users)
	jsonData, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(jsonData))
	f, err := os.OpenFile("test.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	_, err = f.Write(jsonData)
	if err != nil {
		log.Println(err)
		return
	}

}
