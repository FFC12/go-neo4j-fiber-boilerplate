package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(data map[string]any, hour int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(hour) * time.Hour).Unix()

	for key, value := range data {
		InfoLogger.Println("Key: ", key, " - Value: ", value)
		claims[key] = value
	}

	ErrorLogger.Println("Secret: ", JWT_SECRET_KEY)

	tokenStr, err := token.SignedString([]byte(JWT_SECRET_KEY))

	return tokenStr, err
}

func extractClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("failed to extract payload of JWT")
	}

}

func VerifyJWT(tokenString string) map[string]interface{} {
	payload, err := extractClaims(tokenString)

	if err != nil {
		ErrorLogger.Fatal(err)
	}

	return payload
}

func RefreshJWT(data map[string]any, hour int64) (string, error) {
	return GenerateJWT(data, hour)
}
