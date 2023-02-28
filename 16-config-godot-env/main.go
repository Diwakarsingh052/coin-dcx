package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//load the config file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
	//fetch key value
	db := os.Getenv("Database")
	log.Println(db)

	//https://github.com/spf13/viper
	
}
