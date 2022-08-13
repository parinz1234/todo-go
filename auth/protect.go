package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func Protect(tokenString string) error {
	// HOF
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check token method mecanism
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// if valid will pass signature to jwt.Parse as param
		return []byte("==signature=="), nil
	})

	return err
}
