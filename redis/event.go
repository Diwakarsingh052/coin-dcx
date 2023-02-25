package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

func main() {
	log.Println("connecting to redis....")
	var rdb *redis.Client

	rdb, err := ConnectToRedis()

	if err != nil {
		log.Println("all attempts failed , cannot connect to redis", err)
		return
	}

	log.Println("redis connected")

	http.HandleFunc("/ping", ping)

	go getArticleEvent(rdb)

	http.ListenAndServe(":8080", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "notification service is running")
}

func getArticleEvent(r *redis.Client) {
	// for loop will make sure we keep listening for new events
	for {
		// xread in redis allows us to listen for new messages that are not processed
		val := redis.XReadArgs{
			Streams: []string{"article:add:events", "$"}, // article:add:event is key name, $ is used to fetch msg after the last delivered id
			Count:   1,                                   // return 1 event at a time
			Block:   0,                                   // it means wait until there is no new event // unlimited time
		}
		//exec the xread command in redis and waiting for the result
		res, err := r.XRead(context.Background(), &val).Result()

		if err != nil {
			log.Println(err)
		}

		//when new event is read from redis, we will send the notification to all the followers that a new post is out
		go sendNotification(res) // running it as go routine, so we can move forward to listen for the new messages

	}

}

// sendNotification sends a notification to all the followers of the publisher
func sendNotification(res []redis.XStream) {

	//ranging over xread result
	for _, r := range res {
		fmt.Println(r.Stream, "key name")

		for _, msg := range r.Messages {
			fmt.Println(msg.ID, "stream id")
			fmt.Println(msg.Values, "printing values of the map")

			//taking the email out of the article publisher from the event
			v, ok := msg.Values["email"]
			if !ok {
				log.Println("email not found, can't send a notification")
				continue
			}
			//making sure email is of correct type using type assertion
			email, ok := v.(string)
			if !ok {
				log.Println("email type can't be identified, can't send a notification")
				continue
			}

			title, ok := msg.Values["title"]
			if !ok {
				log.Println("title not found, can't send a notification")
				continue
			}

			followersList, ok := followers[email]
			if !ok {
				log.Println("no followers found")
				continue
			}

			//send an email notification to all the followers in a goroutine
			go sendEmail(followersList, title)

		}

	}

}

var followers map[string][]string = map[string][]string{
	"johndoe@example.com": {
		"user1@example.com",
		"user2@example.com",
		"user3@example.com",
		"user4@example.com",
		"user5@example.com",
	},
}

func sendEmail(list []string, title interface{}) {
	for _, e := range list {
		fmt.Println("sending email", e)

		fmt.Println("article title", title)
	}

}
