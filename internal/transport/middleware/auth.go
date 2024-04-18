package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/sail3/interfell-vaccinations/internal/response"
)

const userKey string = "userID"

func Authenticator(signingString string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := validateToken(r, signingString)
			if err != nil {
				response.ResponsdWithData(w, http.StatusForbidden, err)
				return
			}
			ctx := context.WithValue(r.Context(), userKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func validateToken(r *http.Request, signingString string) (id string, err error) {
	auth := r.Header.Get("Authorization")
	tokenString, err := tokenFromAuthorization(auth)
	if err != nil {
		return "", errors.New("access error")
	}
	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return []byte(signingString), nil
	})
	if err != nil {
		return "", errors.New("access error")
	}

	if token == nil || !token.Valid {
		return "", errors.New("access error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("access error")
	}

	id, ok = claims["id"].(string)
	if !ok {
		return "", errors.New("access error")
	}

	return id, nil
}

func tokenFromAuthorization(authorization string) (string, error) {
	if authorization == "" {
		return "", errors.New("access error")
	}

	if !strings.HasPrefix(authorization, "Bearer") {
		return "", errors.New("access error")
	}

	l := strings.Split(authorization, " ")
	if len(l) != 2 {
		return "", errors.New("access error")
	}

	return l[1], nil
}
