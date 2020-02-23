package utils

import (
	"LionChallenge/model"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// GenerateToken function use to generate token
func GenerateToken(user *model.User) (string, error) {
	secretKey := viper.GetString("token.key")

	claims := model.TokenClaims{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// GetDataFromToken function
func GetDataFromToken(request *http.Request) *model.TokenClaims {
	authHeader := request.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	authToken := bearerToken[1]

	token, _ := jwt.ParseWithClaims(authToken, &model.TokenClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(viper.GetString("token.key")), nil
	})
	claims := token.Claims.(*model.TokenClaims)
	return claims
}
