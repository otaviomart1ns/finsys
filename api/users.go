package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/otaviomart1ns/finsys/db/sqlc"
)

type addUserRequest struct {
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	LastName string    `json:"last_name" binding:"required"`
	Birth    time.Time `json:"birth" binding:"required"`
	Email    string    `json:"email" binding:"required"`
}

func (server *Server) addUser(ctx *gin.Context) {
	var req addUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	/* hashedInput := sha512.Sum512_256([]byte(req.Password))
	trimmedHash := bytes.Trim(hashedInput[:], "\x00")
	preparedPassword := string(trimmedHash)
	passwordHashInBytes, err := bcrypt.GenerateFromPassword([]byte(preparedPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	var passwordHashed = string(passwordHashInBytes) */
	params := db.AddUserParams{
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		LastName: req.LastName,
		Birth:    req.Birth,
		Email:    req.Email,
	}

	user, err := server.store.AddUser(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, user)
}

type updateUserRequest struct {
	ID       int32     `json:"id" binding:"required"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Birth    time.Time `json:"birth"`
	Email    string    `json:"email"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	params := db.UpdateUserParams{
		ID:       req.ID,
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		LastName: req.LastName,
		Birth:    req.Birth,
		Email:    req.Email,
	}

	category, err := server.store.UpdateUser(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, category)
}

func (server *Server) getUsers(ctx *gin.Context) {
	user, err := server.store.GetUsers(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type deleteUserRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.DeleteUser(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, true)
}

type getUserByIDRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getUserByID(ctx *gin.Context) {
	var req getUserByIDRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	user, err := server.store.GetUserByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserByEmailRequest struct {
	Email string `uri:"email" binding:"required"`
}

func (server *Server) getUserByEmail(ctx *gin.Context) {
	var req getUserByEmailRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserByUsernameRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) getUserByUsername(ctx *gin.Context) {
	var req getUserByUsernameRequest
	err := ctx.ShouldBindUri(&req)
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

	ctx.JSON(http.StatusOK, user)
}

type getUserByNameAndLastNameRequest struct {
	Name     string `uri:"name" binding:"required"`
	LastName string `uri:"last_name" binding:"required"`
}

func (server *Server) getUserByNameAndLastName(ctx *gin.Context) {
	var req getUserByNameAndLastNameRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	params := db.GetUserByNameAndLastNameParams{
		Name:     req.Name,
		LastName: req.LastName,
	}

	user, err := server.store.GetUserByNameAndLastName(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserByEmailAndPasswordRequest struct {
	Email    string `uri:"email" binding:"required"`
	Password string `uri:"password" binding:"required"`
}

func (server *Server) getUserByEmailAndPassword(ctx *gin.Context) {
	var req getUserByEmailAndPasswordRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	params := db.GetUserByEmailAndPasswordParams{
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := server.store.GetUserByEmailAndPassword(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
