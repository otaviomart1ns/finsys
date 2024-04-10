package api

import (
	"bytes"
	"crypto/sha512"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	user, err := server.store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	hashedInput := sha512.Sum512_256([]byte(req.Password))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPassword := string(trimmedHash)

	passwordInBytes := []byte(preparedPassword)
	passwordHashInBytes := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(passwordHashInBytes, passwordInBytes)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	params := &loginResponse{
		UserID: user.ID,
		Token:  tokenString,
	}

	ctx.JSON(http.StatusOK, params)
}
