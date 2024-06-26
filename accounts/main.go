package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	commonDB "github.com/otaviomart1ns/finsys/common/db/sqlc"
	"github.com/otaviomart1ns/finsys/common/utils"
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

func AddAccount(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	var req addAccountRequest
	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	var categoryId = req.CategoryID
	var accountType = req.Type

	category, err := store.GetCategoryByID(ctx, categoryId)
	if err != nil {
		return utils.ErrorResponse(http.StatusNotFound, fmt.Errorf("error searching category by ID: %v", err)), nil
	}

	var categoryTypeDiffAccountType = category.Type != accountType
	if categoryTypeDiffAccountType {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	} else {
		params := commonDB.AddAccountParams{
			UserID:      req.UserID,
			CategoryID:  categoryId,
			Title:       req.Title,
			Type:        accountType,
			Description: req.Description,
			Value:       req.Value,
			Date:        req.Date,
		}

		account, err := store.AddAccount(ctx, params)
		if err != nil {
			return utils.ErrorResponse(http.StatusInternalServerError, err), nil
		}

		return utils.Response(http.StatusOK, account)
	}

}

type updateAccountRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
}

func UpdateAccount(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	var req updateAccountRequest
	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	params := commonDB.UpdateAccountParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Value:       req.Value,
	}

	account, err := store.UpdateAccount(ctx, params)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}

	return utils.Response(http.StatusOK, account)
}

func DeleteAccount(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	accountID, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid account ID")), nil
	}

	err = store.DeleteAccount(ctx, int32(accountID))
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, fmt.Errorf("error deleting account: %v", err)), nil
	}

	message := fmt.Sprintf("Account with ID %d successfully deleted.", accountID)
	return utils.Response(http.StatusOK, message)
}

func GetAccounts(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	reqUserID, err := strconv.Atoi(request.QueryStringParameters["user_id"])
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err)), nil
	}
	reqType := request.QueryStringParameters["type"]
	reqTitle := request.QueryStringParameters["title"]
	reqDescription := request.QueryStringParameters["description"]
	reqCategoryID, _ := strconv.Atoi(request.QueryStringParameters["category_id"])
	reqDate, _ := time.Parse(time.RFC3339, request.QueryStringParameters["date"])

	params := commonDB.GetAccountsParams{
		UserID:      int32(reqUserID),
		Type:        reqType,
		Title:       reqTitle,
		Description: reqDescription,
		CategoryID: sql.NullInt32{
			Int32: int32(reqCategoryID),
			Valid: int32(reqCategoryID) > 0,
		},
		Date: sql.NullTime{
			Time:  reqDate,
			Valid: !reqDate.IsZero(),
		},
	}

	accounts, err := store.GetAccounts(ctx, params)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	return utils.Response(http.StatusOK, accounts)
}

func GetAccountByID(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	accountID, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid account ID")), nil
	}

	account, err := store.GetAccountByID(ctx, int32(accountID))
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, fmt.Errorf("error searching account by ID: %v", err)), nil
	}

	return utils.Response(http.StatusOK, account)
}

func GetAccountGraph(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	userID, err := strconv.Atoi(request.PathParameters["user_id"])
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid user ID")), nil
	}

	typeReq := request.PathParameters["type"]

	params := commonDB.GetAccountGraphParams{
		UserID: int32(userID),
		Type:   typeReq,
	}

	accountGraph, err := store.GetAccountGraph(ctx, params)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}

	return utils.Response(http.StatusOK, accountGraph)
}

func GetAccountReports(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	userID, err := strconv.Atoi(request.PathParameters["user_id"])
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid user ID")), nil
	}

	typeReq := request.PathParameters["type"]

	params := commonDB.GetAccountReportsParams{
		UserID: int32(userID),
		Type:   typeReq,
	}

	accountReports, err := store.GetAccountReports(ctx, params)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}

	return utils.Response(http.StatusOK, accountReports)
}
