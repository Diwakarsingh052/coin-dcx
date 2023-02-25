package main

import (
	pb "client/gen/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"

	"time"
)

func main() {
	log := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStreamInterceptor(clientStreamInterceptor),
	}
	conn, err := grpc.Dial("localhost:5000", dialOpts...)
	//conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)

	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	//go func() {
	//	time.Sleep(time.Second)
	//	cancel()
	//}()

	stream, err := client.CreatePost(ctx)
	if err != nil {
		log.Fatalf("failed to call createPost server: %v", err)
	}

	//assume these posts we are getting in batches
	batch1 := []*pb.Post{
		{
			Title:  "The Science of Design",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "The Politics of Power",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "The Art of Programming",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}
	p := &pb.CreatePostRequest{Posts: batch1}

	err = stream.Send(p)
	if err != nil {
		log.Fatalf("Failed to createPost request: %v", err)
	}

	//adding latency
	time.Sleep(4 * time.Second)

	//constructing the second batch
	batch2 := []*pb.Post{
		{
			Title:  "Post 11",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "Post 21",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "Post 31",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}
	p = &pb.CreatePostRequest{Posts: batch2}

	//sending the second batch to the server to process
	//if the context was cancelled already then calling of this method would fail
	err = stream.Send(p)
	if err != nil {
		log.Fatalf("Failed to createPost request: %v", err)
	}

	//close a client-streaming, and receive the server's response message.
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}

	log.Printf("Response: %s", response.Result)
}
