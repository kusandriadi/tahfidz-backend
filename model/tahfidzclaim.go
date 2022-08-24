package model

import "github.com/golang-jwt/jwt/v4"

type TahfidzClaim struct {
	jwt.RegisteredClaims
	Id       int    `json:"Id"`
	Username string `json:"Username"`
	Role     string `json:"Role"`
}
