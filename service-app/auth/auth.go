package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

// Two roles
// Admin // User
const (
	RoleAdmin = "ADMIN"
	RoleUser  = "USER"
)

// ctxKey type would be used to put the claims in the context
type ctxKey int

const Key ctxKey = 1

// Claims is our payload/data for out jwt token
type Claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

// HasRoles checks if the data has a valid role to access the endpoint
func (c Claims) HasRoles(requiredRoles ...string) bool {
	for _, has := range c.Roles { // roles with the users
		for _, want := range requiredRoles { // roles data required to access the endpoint
			if has == want {
				return true
			}
		}
	}

	return false
}

// Auth struct privateKey field would be used to verify and generate token
type Auth struct {
	privateKey *rsa.PrivateKey // this is key we would get after parsing the private.pem file
	publicKey  *rsa.PublicKey
}

// NewAuth func set the privateKey in the Auth struct and returns the instance of it to the caller
func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*Auth, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("private/public key cannot be nil")
	}

	a := Auth{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
	return &a, nil
}

func (a *Auth) GenerateToken(claims Claims) (string, error) {
	//jwt.NewWithClaims takes a signingMethod and claims struct to generate a token on basis of it
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, &claims)

	//signing our token with our private key
	tokenStr, err := tkn.SignedString(a.privateKey)

	if err != nil {
		return "", fmt.Errorf("signing token %w", err)
	}

	return tokenStr, nil

}

func (a *Auth) ValidateToken(tokenStr string) (Claims, error) {
	var claims Claims

	//jwt.ParseWithClaims verify the token against the public key
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})

	if err != nil {
		return Claims{}, fmt.Errorf("parsing token %w", err)
	}
	if !token.Valid {
		return Claims{}, errors.New("invalid token")
	}

	//returning Claims if verification is successful
	return claims, nil

}
