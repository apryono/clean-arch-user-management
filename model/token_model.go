package model

import "github.com/dgrijalva/jwt-go"

// TokenClaims for jwt
type TokenClaims struct {
	jwt.StandardClaims
	UserID   uint64 `from:"UserID" json:"user_id"`
	Username string `from:"Username" json:"username"`
	Email    string `from:"Email" json:"email"`
}
