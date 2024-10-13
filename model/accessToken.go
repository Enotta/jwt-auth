package model

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// Access token model
type AccessToken struct {
}

func CreateAccessToken(user User) string {
	var key []byte = []byte("secretSecretsecret")

	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"iss": "auth",
			"sub": user.Name + " " + user.Email,
			"foo": fmt.Sprint(user.ID),
		})

	s, err := token.SignedString(key)
	if err != nil {
		fmt.Println("'" + s + "'")
		return s
	}

	panic("Failed to create access token")
}
