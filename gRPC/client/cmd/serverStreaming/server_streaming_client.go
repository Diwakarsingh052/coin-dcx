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
)

func main() {

	log := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)

	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	req := &pb.GetPostsRequest{UserId: 101}

	// doing server streaming request
	stream, err := client.GetPosts(context.Background(), req)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		//receiving values from stream
		post, err := stream.Recv()

		//if the server has finished sending the request, we will quit
		if err == io.EOF {
			break
		}
		//any other kind of error would be caught here
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("reading stream")

		//printing data received
		fmt.Println(post)
		fmt.Println()
	}
}
