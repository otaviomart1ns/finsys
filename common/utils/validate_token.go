package utils

import (
	"errors"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ValidateToken(token string) (*Claims, error) {
	claims := &Claims{}
	var jwtSignedKey = []byte("secret_key")
	tokenParse, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtSignedKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("unauthorized: invalid signature")
		}
		return nil, err
	}

	if !tokenParse.Valid {
		return nil, errors.New("unauthorized: token is invalid")
	}

	return claims, nil
}

func GetTokenAndVerify(request events.APIGatewayProxyRequest) (*Claims, error) {
	authHeader := request.Headers["Authorization"]
	if authHeader == "" {
		return nil, errors.New("unauthorized: no authorization header provided")
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, errors.New("unauthorized: invalid authorization header format")
	}

	token := fields[1]
	return ValidateToken(token)
}
