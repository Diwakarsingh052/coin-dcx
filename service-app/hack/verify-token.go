package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
)

const tokenStr = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEwMSIsImV4cCI6MTY3NjUyNjM2NCwiaWF0IjoxNjc2NTIzMzY0LCJyb2xlcyI6WyJVU0VSIl19.kBFGbaWG5mlVZVt9TdWfpKoA3PGDQyPqhptg41JIJTMU-8BEWynb4swzDaOcQPnEliF19ctkOylo1KHKjhR5P5y8-01Df_jGpjq_F-i3_pC9KP9vp4bWeycx_nzAtX-YpQmhJo2I_khBg_nk2rBtKz7zugj2l2lEGZ6r90ZlhT1EfozKNwTL2yeskpgINxFhmXltIUvPF6_abFYBjV7ljDNzJhZOudiAmmCvQVhZvIvz_Ck1TRSoDBdChNLeii6u_90wcug7wMgae8IKJZlDmdfTGahcAeYq3YJAJTXe_wf7RhAHPqO1WX7PkiLjnZVTNi8Svv9jqOO_-5d9RdfNtw`

func main() {

	type claims struct {
		jwt.RegisteredClaims
		Roles []string `json:"roles"`
	}

	PublicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		log.Fatalln("not able to read pem file")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicPEM)

	if err != nil {
		log.Fatalln(err)
	}

	var c claims // claims struct will hold jwt data after parsing the token

	//jwt.ParseWithClaims verify the token against the public key
	token, err := jwt.ParseWithClaims(tokenStr, &c, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		fmt.Println("parsing token", err)
		return
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return
	}

	fmt.Println(c)
}
