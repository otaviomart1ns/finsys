package main

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	commonDB "github.com/otaviomart1ns/finsys/common/db/sqlc"
	"github.com/otaviomart1ns/finsys/common/utils"
	"golang.org/x/crypto/bcrypt"
)

type addUserRequest struct {
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	LastName string    `json:"last_name" binding:"required"`
	Birth    time.Time `json:"date" binding:"required"`
	Email    string    `json:"email" binding:"required"`
}

func AddUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req addUserRequest
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	hashedInput := sha512.Sum512_256([]byte(req.Password))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPassword := string(trimmedHash)
	passwordHashInBytes, err := bcrypt.GenerateFromPassword([]byte(preparedPassword), bcrypt.DefaultCost)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}
	passwordHashed := string(passwordHashInBytes)

	params := commonDB.AddUserParams{
		Username: req.Username,
		Password: passwordHashed,
		Name:     req.Name,
		LastName: req.LastName,
		Birth:    req.Birth,
		Email:    req.Email,
	}

	user, err := store.AddUser(ctx, params)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}

	return utils.Response(http.StatusOK, user)
}

type updateUserRequest struct {
	ID          int32     `json:"id" binding:"required"`
	Username    string    `json:"username"`
	OldPassword string    `json:"old_password"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	Birth       time.Time `json:"birth"`
	Email       string    `json:"email"`
}

func UpdateUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	var req updateUserRequest
	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	user, err := store.GetUserByID(ctx, req.ID)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}

	var passwordHash string = user.Password

	if req.Password != "" {
		hashedOldInput := sha512.Sum512_256([]byte(req.OldPassword))
		trimmedOldHash := bytes.Trim(hashedOldInput[:], "\x00")
		preparedOldPassword := string(trimmedOldHash)

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(preparedOldPassword)); err != nil {
			return utils.ErrorResponse(http.StatusUnauthorized, err), nil
		}

		hashedInput := sha512.Sum512_256([]byte(req.Password))
		trimmedHash := bytes.Trim(hashedInput[:], "\x00")
		preparedPassword := string(trimmedHash)

		passwordHashInBytes, err := bcrypt.GenerateFromPassword([]byte(preparedPassword), bcrypt.DefaultCost)
		if err != nil {
			return utils.ErrorResponse(http.StatusInternalServerError, err), nil

		}
		passwordHash = string(passwordHashInBytes)
	} else {
		passwordHash = user.Password
	}

	params := commonDB.UpdateUserParams{
		ID:       req.ID,
		Username: utils.Coalesce(req.Username, user.Username),
		Password: passwordHash,
		Name:     utils.Coalesce(req.Name, user.Name),
		LastName: utils.Coalesce(req.LastName, user.LastName),
		Birth:    utils.CoalesceTime(req.Birth, user.Birth),
		Email:    utils.Coalesce(req.Email, user.Email),
	}

	updateUser, err := store.UpdateUser(ctx, params)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}

	return utils.Response(http.StatusOK, updateUser)
}

func DeleteUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userID, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid user ID")), nil
	}

	err = store.DeleteUser(ctx, int32(userID))
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, fmt.Errorf("error deleting user: %v", err)), nil
	}

	message := fmt.Sprintf("User with ID %d successfully deleted.", userID)
	return utils.Response(http.StatusOK, message)
}

func GetUsers(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	users, err := store.GetUsers(ctx)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	if len(users) == 0 {
		return utils.Response(http.StatusNotFound, []interface{}{})
	}

	return utils.Response(http.StatusOK, users)
}

func GetUserByID(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userID, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid user ID")), nil
	}

	user, err := store.GetUserByID(ctx, int32(userID))
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, fmt.Errorf("error searching user by ID: %v", err)), nil
	}

	return utils.Response(http.StatusOK, user)
}

func GetUserByUsername(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	username, ok := request.PathParameters["username"]
	if !ok || username == "" {
		return utils.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid username")), nil
	}

	user, err := store.GetUserByUsername(ctx, username)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, fmt.Errorf("error searching user by username: %v", err)), nil
	}

	return utils.Response(http.StatusOK, user)
}
