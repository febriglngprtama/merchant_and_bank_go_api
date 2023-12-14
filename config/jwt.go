package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("720bd276f6c5b9dcceba74056a49661c")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}