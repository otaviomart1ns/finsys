package main

import (
	"bytes"
	"context"
	"crypto/sha512"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"
	"github.com/otaviomart1ns/finsys/common/utils"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	UserID int32  `json:"user_id"`
	Token  string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req loginRequest
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	user, err := store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return utils.ErrorResponse(http.StatusNotFound, err), nil
		}
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}

	hashedInput := sha512.Sum512_256([]byte(req.Password))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPassword := string(trimmedHash)

	passwordInBytes := []byte(preparedPassword)
	passwordHashInBytes := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(passwordHashInBytes, passwordInBytes)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	expirationToken := time.Now().Add(time.Hour * 24)

	claims := &Claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationToken),
		},
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtSignedKey = []byte("secret_key")
	tokenString, err := generateToken.SignedString(jwtSignedKey)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}

	params := &loginResponse{
		UserID: user.ID,
		Token:  tokenString,
	}

	return utils.Response(http.StatusOK, params)
}
