package token

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

var SampleSecretKey = []byte("veryverysecretverysecret")

type CustomClaims struct {
	UserId uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJwt(userId uuid.UUID, email string) (string, error) {

	claims := &CustomClaims{
		userId,
		email,
		jwt.RegisteredClaims{
			Issuer:    "space",
			Subject:   "",
			Audience:  jwt.ClaimStrings{"client"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//claims := token.Claims.(jwt.MapClaims)
	//claims["exp"] = time.Now().UTC().Add(24 * time.Hour)
	//claims["nvb"] = time.Now().UTC()
	//claims["iat"] = time.Now().UTC()
	//claims["iss"] = "space"
	//claims["aud"] = "client"
	//for key, claim := range customClaims {
	//	claims[key] = claim
	//}

	tokenString, err := token.SignedString(SampleSecretKey)
	return tokenString, err
}
