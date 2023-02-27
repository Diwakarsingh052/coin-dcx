package main

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	pb "server/gen/proto"

	"strings"

	"github.com/golang-jwt/jwt/v4"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

//https://github.com/grpc-ecosystem/go-grpc-middleware

var pubKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyqwTw0Q11D7+FWrSW6T0
vdka5ZvKfxWvEQlyXpRLCQFOhzUU5mey2dievHU6TyO0WrH9Y7/mHnHTAuTK1zrK
e62EjgR8cIbs0ikvXWq0A/8zfo5SbXIRTANM5yRl35OBAn/FUhKXeR/GIuSpDmgk
CD74cPoskQODfQfjuGT+D6A7uEFnhhka4unwyKcYSrAhIpitwWyqqcQEwJpUupby
oM6WijVYZt87U+Es/U7weI3N0PBAG9A1fcSZ+BifaVNT1EaO+ETpJPnMDmbe2Pfg
NGdjbNDDQLPich9dboZPSVVqW3OSyJFkJoW9jOS0SQmD+2pgNMA922OQUiwt3DKF
gwIDAQAB
-----END PUBLIC KEY-----`

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Jwt(ctx context.Context, req *pb.JwtRequest) (*pb.JwtResponse, error) {
	name := req.GetName()
	fmt.Println(ctx.Value("data"))
	response := &pb.JwtResponse{Result: "Hello, " + name}
	return response, nil
}

func jwtAuthFunc(ctx context.Context) (context.Context, error) {
	//parsing public key
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKey))
	if err != nil {
		log.Println("public key invalid")
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	// Get the metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	// Get the JWT token from the metadata
	authHeader := md.Get("authorization")
	fmt.Println(authHeader)

	//making sure we have received the value in metadata
	if len(authHeader) == 0 {

		log.Printf("%#v", "expected authorization header format: Bearer <token>")
		return nil, status.Error(codes.Unauthenticated, "expected authorization header format: Bearer <token>")
	}

	// Remove the "Bearer " prefix from the token string
	tokenString := strings.Split(authHeader[0], " ")

	//making sure token is provided in correct format
	if strings.ToLower(tokenString[0]) != "bearer" {
		log.Printf("%#v", "expected authorization header format: Bearer <token>")
		return nil, status.Error(codes.Unauthenticated, "expected authorization header format: Bearer <token>")

	}
	//fmt.Println(tokenString[0])
	//printing token
	//fmt.Println(tokenString[1])

	//taking token value out of the slice
	token := tokenString[1]

	// Parse and validate the JWT token
	claims, err := ValidateToken(token, publicKey)
	if err != nil {
		log.Println("token invalid")
		return nil, status.Error(codes.Unauthenticated, "token invalid")
	}

	//data have valid token
	// Add the data info to the context // it will tell us who is logged-in
	// In this example, we're adding the username at the "data" key
	username := claims.Subject
	return context.WithValue(ctx, "data", username), nil
}

func ValidateToken(tokenStr string, pubKey *rsa.PublicKey) (jwt.RegisteredClaims, error) {
	var claims jwt.RegisteredClaims

	// verifying token with our public key // if token is valid, we fetch the data stored inside the token and put the data in claims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})

	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("parsing token %w", err)
	}

	if !token.Valid {
		return jwt.RegisteredClaims{}, errors.New("invalid token")
	}

	return claims, nil

}

func main() {
	log := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
	// Add the JWT middleware to the gRPC server

	//UnaryServerInterceptor returns a new unary server interceptors
	//that performs per-request auth
	authUnaryInterceptor := grpc_auth.UnaryServerInterceptor(jwtAuthFunc)
	serverOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(authUnaryInterceptor),
	}
	// Create a new gRPC server with options
	s := grpc.NewServer(serverOpts...)

	// Register your gRPC server implementation
	pb.RegisterUserServiceServer(s, &server{})

	// Start the gRPC server
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
