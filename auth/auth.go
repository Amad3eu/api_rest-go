package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		authStr := strings.Split(authHeader, " ")
		if len(authStr) == 2 && authStr[0] == "Basic" {
			err := BasicAuth(req)
			if err != nil {
				http.Error(res, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(res, req)
		} else if len(authStr) == 2 && authStr[0] == "Bearer" {
			err := JwtAuthentication(authHeader)
			if err != nil {
				http.Error(res, "Access token is missing or invalid", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(res, req)
		} else {
			http.Error(res, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func BasicAuth(req *http.Request) error {
	username, password, ok := req.BasicAuth()
	if !ok {
		return errors.New("invalid authentication header")
	}
	if username != os.Getenv("ADMIN_USER") || password != os.Getenv("ADMIN_PASS") {
		return errors.New("wrong username or password")
	}
	return nil
}

func GenerateToken(userID string) (string, error) {
	signingKey := []byte(os.Getenv("SIGNING_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		Id:        userID,
		Issuer:    "Admin",
	})
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func JwtAuthentication(authHeader string) error {
	authStr := strings.Split(authHeader, " ")
	tokenString := authStr[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error")
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return err
	} else if _, ok := err.(*jwt.ValidationError); ok {
		return errors.New("access token is missing or invalid")
	} else if token.Valid {
		return nil
	} else {
		return errors.New("access token is missing or invalid")
	}
}
