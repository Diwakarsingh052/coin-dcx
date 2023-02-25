package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
)

func main() {

	log := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	// establishes a gRPC client connection to a server
	conn, err := grpc.Dial("localhost:5000", dialOpts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	//creates a gRPC client for the UserService service defined in the pb package
	//and binds it to the specified connection conn.
	client := pb.NewUserServiceClient(conn)

	//----- Sending Metadata to server-----//

	//metadata.Pairs to create metadata in pairs,
	//metadata with the same key will get merged into a list:
	md := metadata.Pairs(
		"test", "metadata",
		"key", "value",
		"key", "value2",
	)

	//NewOutgoingContext creates a new context with outgoing md attached
	//NewOutgoingContext will overwrite any previously-appended metadata.
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)

	//AppendToOutgoingContext returns a new context with the provided kv merged with any existing metadata in the context
	ctxA := metadata.AppendToOutgoingContext(mdCtx,
		"k1", "v1", "k1", "v2", "k2", "v3")

	//----- Recv Metadata from server-----//
	var header, trailer metadata.MD
	// make unary RPC
	response, err := client.Hello(
		ctxA,
		&pb.HelloRequest{Name: "john"},
		grpc.Header(&header),   //Header returns a CallOptions that retrieves the header metadata for a unary RPC.
		grpc.Trailer(&trailer), //Trailer returns a CallOptions that retrieves the trailer metadata for a unary RPC
	)
	if err != nil {
		log.Println(err)
		return
	}
	// process header and trailer map here
	fmt.Println(header, trailer)
	fmt.Println(response)

}
