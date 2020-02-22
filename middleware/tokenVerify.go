package middleware

import (
	"LionChallenge/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// TokenVerifyMiddleware to verify token
func TokenVerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (i interface{}, err error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There Was An Error %w", err)
				}
				return []byte(viper.GetString("token.key")), nil
			})

			if err != nil {
				fmt.Errorf("There are error with : %w", err)
				utils.Response(w, utils.Message(false, "Oops.. Something went wrong"))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				utils.Response(w, utils.Message(false, "Invalid Token"))
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

		} else {
			utils.Response(w, utils.Message(false, "Invalid Token"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}
