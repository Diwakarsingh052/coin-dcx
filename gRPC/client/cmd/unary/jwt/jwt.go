package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	log := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a new gRPC client using the connection
	client := pb.NewUserServiceClient(conn)

	//generating token
	signedToken, err := genToken()
	if err != nil {
		log.Fatalf("Failed to sign token: %v", err)
	}

	// Add the JWT token to the gRPC metadata for authentication
	md := metadata.Pairs("authorization", "bearer "+signedToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// Make a gRPC request with the authenticated context
	resp, err := client.Jwt(ctx, &pb.JwtRequest{Name: "John"})
	if err != nil {
		log.Fatalf("error in Jwt rpc : %v", err)
	}

	// Handle the gRPC response
	fmt.Printf("Response: %v\n", resp)
}

func genToken() (string, error) {
	PrivatePem, err := os.ReadFile("private.pem")
	if err != nil {
		return "", err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(PrivatePem)
	if err != nil {
		return "", err
	}
	// iss (issuer): Issuer of the JWT
	// sub (subject): Subject of the JWT (the users)
	// aud (audience): Recipient for which the JWT is intended
	// exp (expiration time): Time after which the JWT expires
	// nbf (not before time): Time before which the JWT must not be accepted for processing
	// iat (issued at time): Time at which the JWT was issued; can be used to determine age of the JWT
	// jti (JWT ID): Unique identifier; can be used to prevent the JWT from being replayed (allows a token to be used only once)

	claims := jwt.RegisteredClaims{
		Issuer:    "api project",
		Subject:   "john@email.com",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(50 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	str, err := tkn.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return str, nil
}
