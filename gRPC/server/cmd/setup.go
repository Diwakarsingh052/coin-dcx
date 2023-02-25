package main

import (
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"server/database"
	pb "server/gen/proto"
)

type userService struct {
	pb.UnimplementedUserServiceServer
	db  *sql.DB
	log *log.Logger
}

func main() {
	log := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
	//sets up a TCP listener on port 5000 for incoming network connections.
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Println(err)
		return
	}

	//NewServer creates a gRPC server which has no service registered
	//and has not started to accept requests yet.

	srvOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(UnaryServerInterceptor),
		grpc.StreamInterceptor(ServerStreamInterceptor),
	}
	s := grpc.NewServer(srvOptions...)
	//s := grpc.NewServer()
	//connecting to database
	db, err := database.ConnectToPostgres()
	if err != nil {
		log.Fatalln(err)
	}

	//registers a gRPC service implementation with a gRPC server.
	//The second argument is a pointer to the service implementation struct,
	//which implements the methods defined in the gRPC service interface.
	pb.RegisterUserServiceServer(s, &userService{db: db, log: log})

	//exposing gRPC service to be tested by postman
	reflection.Register(s)
	//Serve accepts incoming connections on the listener lis
	if err := s.Serve(listener); err != nil {
		fmt.Println(err)
		return
	}

}
