package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	ID string `json:"id" bson:"_id"`
	jwt.RegisteredClaims
}

func GenerateJWToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

func VerifyJWToken(token string) (*UserClaims, error) {
	claims := &UserClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	if !tkn.Valid {
		return nil, err
	}
	return claims, nil
}
