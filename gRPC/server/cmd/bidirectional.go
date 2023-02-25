package main

import (
	"fmt"
	"io"
	pb "server/gen/proto"
)

//In bidirectional-streaming RPC, the client is sending a request to the server as a stream of messages.
//The server also responds with a stream of messages.
//The call has to be initiated from the client side,
//But after that,the communication is completely based on the application logic of the gRPC client and the server.

func (us *userService) GreetEveryone(stream pb.UserService_GreetEveryoneServer) error {
	us.log.Println("GreetEveryone was invoked")

	for {
		//receiving the streaming request from the client
		req, err := stream.Recv()

		//If the client has finished sending the request, we will quit
		if err == io.EOF {
			return nil
		}

		if err != nil {
			us.log.Fatalf("error while reading client stream: %v\n", err)
		}
		us.log.Println("client sent a name to sat hello", req.FirstName)
		fmt.Println()
		res := "Hello " + req.FirstName + "!"

		err = stream.Send(&pb.GreetEveryoneResponse{
			Result: res,
		})

		if err != nil {
			us.log.Fatalf("error while sending data to client: %v\n", err)
		}
	}
}
