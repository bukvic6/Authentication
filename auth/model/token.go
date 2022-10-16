package model

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Email    string
	Password string
	*jwt.StandardClaims
}
