package main

import (
	"context"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "server/gen/proto"
)

type User struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Roles        pq.StringArray `json:"roles"`
	PasswordHash string         `json:"-"`
}

func (us *userService) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	us.log.Println(req.GetUser())
	nu := req.GetUser() // fetching the request sent by the client

	u := User{
		Name:  nu.Name,
		Email: nu.Email,
		Roles: nu.Roles,
	}

	//bcrypt generate hash from the password provided
	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		us.log.Println(err)
		// Custom error with status code and message
		return nil, status.Newf(codes.Internal, "generating password hash %v", err).Err()
	}

	u.PasswordHash = string(hash)

	//constructing query to insert data, and it returns id that is generated by db
	const q = `INSERT INTO users
		(name, email, password_hash, roles)
		VALUES ( $1, $2, $3, $4)
		Returning id`

	//executing the query // QueryRowContext exec the query and returns one row back
	row := us.db.QueryRowContext(ctx, q, u.Name, u.Email, u.PasswordHash, u.Roles)

	//saving the value return by db in id var
	err = row.Scan(&u.ID)

	if err != nil {
		us.log.Println(err)
		// Custom error with status code and message
		return nil, status.Newf(codes.Internal, "account creation failed %v", err).Err()
	}

	us.log.Println(u)

	return &pb.SignupResponse{Result: u.Email + " account created"}, nil
}
