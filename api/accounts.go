package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/otaviomart1ns/finsys/db/sqlc"
	"github.com/otaviomart1ns/finsys/utils"
)

type addAccountRequest struct {
	UserID      int32     `json:"user_id" binding:"required"`
	CategoryID  int32     `json:"category_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Value       int32     `json:"value" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
}

func (server *Server) addAccount(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req addAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	var categoryId = req.CategoryID
	var accountType = req.Type

	category, err := server.store.GetCategoryByID(ctx, categoryId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
	}

	var categoryTypeDiffAccountType = category.Type != accountType
	if categoryTypeDiffAccountType {
		ctx.JSON(http.StatusBadRequest, "Account type is different of Category type")
	} else {
		params := db.AddAccountParams{
			UserID:      req.UserID,
			CategoryID:  categoryId,
			Title:       req.Title,
			Type:        accountType,
			Description: req.Description,
			Value:       req.Value,
			Date:        req.Date,
		}

		account, err := server.store.AddAccount(ctx, params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}

		ctx.JSON(http.StatusOK, account)
	}
}

type updateAccountRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req updateAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	params := db.UpdateAccountParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Value:       req.Value,
	}

	account, err := server.store.UpdateAccount(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, account)
}

type deleteAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req deleteAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, true)
}

type getAccountsRequest struct {
	UserID      int32     `json:"user_id" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CategoryID  int32     `json:"category_id"`
	Date        time.Time `json:"date"`
}

func (server *Server) getAccounts(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req getAccountsRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := db.GetAccountsParams{
		UserID:      req.UserID,
		Type:        req.Type,
		Title:       req.Title,
		Description: req.Description,
		CategoryID: sql.NullInt32{
			Int32: req.CategoryID,
			Valid: req.CategoryID > 0,
		},
		Date: sql.NullTime{
			Time:  req.Date,
			Valid: !req.Date.IsZero(),
		},
	}

	accounts, err := server.store.GetAccounts(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type getAccountByIDRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getAccountByID(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req getAccountByIDRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := server.store.GetAccountByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountGraphRequest struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountGraph(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req getAccountGraphRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	params := db.GetAccountGraphParams{
		UserID: req.UserID,
		Type:   req.Type,
	}

	accountGraph, err := server.store.GetAccountGraph(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accountGraph)
}

type getAccountReportsRequest struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) getAccountReports(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req getAccountReportsRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	params := db.GetAccountReportsParams{
		UserID: req.UserID,
		Type:   req.Type,
	}

	accountReports, err := server.store.GetAccountReports(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accountReports)
}
