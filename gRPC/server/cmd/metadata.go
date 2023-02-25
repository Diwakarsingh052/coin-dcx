package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	pb "server/gen/proto"
)

func (us *userService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := in.GetName()

	//----- Recv Metadata from client-----//
	// Read the metadata map from the incoming context of the remote method / reading metadata sent by client.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("metadata not found")
		return nil, status.Error(codes.Internal, "metadata not found")
	}
	//metadata received from the client
	fmt.Println(md)

	//----- Sending Metadata to client-----//
	// create and send header metadata to client from server
	header := metadata.Pairs("header-key", "val")
	err := grpc.SendHeader(ctx, header)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to send header metadata")
	}

	trailer := metadata.Pairs("trailer-key", "val")
	// create and set trailer
	//SetTrailer sets the trailer metadata that will be sent when an RPC returns.
	err = grpc.SetTrailer(ctx, trailer)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to send trailer metadata")
	}

	return &pb.HelloResponse{Message: name + " metadata found"}, nil
	// do something with metadata

}

//func (us *userService) CreatePost(
//	stream pb.UserService_CreatePostServer) error {
//	//Read the metadata map from the incoming context of the remote method.
//	md, ok := metadata.FromIncomingContext(stream.Context())
//
//	// do something with metadata
//}
