package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ValidateToken(ctx *gin.Context, token string) error {
	claims := &Claims{}
	var jwtSignedKey = []byte("secret_key")
	tokenParse, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtSignedKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, err)
			return err
		}
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}

	if !tokenParse.Valid {
		ctx.JSON(http.StatusUnauthorized, "Token is invalid")
		return nil
	}

	ctx.Next()
	return nil
}

func GetTokenAndVerify(ctx *gin.Context) error {
	authHeaderKey := ctx.GetHeader("authorization")
	fields := strings.Fields(authHeaderKey)
	validateToken := fields[1]
	errValiteToken := ValidateToken(ctx, validateToken)
	if errValiteToken != nil {
		return errValiteToken
	}
	return nil
}
