package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

//https://www.boredapi.com/api/activity
//https://official-joke-api.appspot.com/random_joke

func main() {

	log := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)

	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stream, err := client.GreetEveryone(ctx)

	if err != nil {
		log.Fatalf("failed to call GreetEveryone stream: %v\n", err)
	}
	//requests := make(chan *pb.GreetEveryoneRequest, 10)
	//go func() {
	//	for
	//}()

	requests := []*pb.GreetEveryoneRequest{
		{FirstName: "John"},
		{FirstName: "Bruce"},
		{FirstName: "Roy"},
	}

	//using WaitGroup to wait for goroutines to finish
	wg := &sync.WaitGroup{}
	//send and recv from streams is inside go routines, so they can run independently
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, req := range requests {
			log.Printf("Sending message: %v\n", req)
			//sending requests
			err := stream.Send(req)
			if err != nil {
				log.Println(err)
				return
			}
			time.Sleep(1 * time.Second)
		}
		//closing stream when the client finished sending requests
		err := stream.CloseSend()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Printf("stream has ended")
				break
			}
			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Received: %v\n", res.Result)
		}

	}()

	wg.Wait()
	fmt.Println("end of bidirectional communication")

}
