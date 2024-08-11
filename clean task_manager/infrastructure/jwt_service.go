package infrastructure

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func TokenValidate(auth string) error {
	authSplit := strings.Split(auth, " ")

	if len(authSplit) != 2 || strings.ToLower(authSplit[0]) != "bearer" {
		return errors.New("not authorized")
	}

	token, err := jwt.Parse(authSplit[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return errors.New("invalid jwt")
	}

	return nil
}