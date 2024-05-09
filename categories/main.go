package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	commonDB "github.com/otaviomart1ns/finsys/common/db/sqlc"
	"github.com/otaviomart1ns/finsys/common/utils"
)

type addCategoryRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func AddCategory(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	var req addCategoryRequest
	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	params := commonDB.AddCategoryParams{
		UserID:      req.UserID,
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
	}

	category, err := store.AddCategory(ctx, params)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}

	return utils.Response(http.StatusOK, category)
}

type updateCategoryRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func UpdateCategory(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	var req updateCategoryRequest
	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	params := commonDB.UpdateCategoryParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
	}

	category, err := store.UpdateCategory(ctx, params)
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, err), nil
	}

	return utils.Response(http.StatusOK, category)
}

func DeleteCategory(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	categoryID, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid category ID")), nil
	}

	err = store.DeleteUser(ctx, int32(categoryID))
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, fmt.Errorf("error deleting category: %v", err)), nil
	}

	message := fmt.Sprintf("Category with ID %d successfully deleted.", categoryID)
	return utils.Response(http.StatusOK, message)
}

func GetCategories(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received query parameters: %v", request.QueryStringParameters)

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

	params := commonDB.GetCategoriesParams{
		UserID:      int32(reqUserID),
		Type:        reqType,
		Title:       reqTitle,
		Description: reqDescription,
	}

	categories, err := store.GetCategories(ctx, params)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	return utils.Response(http.StatusOK, categories)
}

func GetCategoryByID(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	categoryID, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, fmt.Errorf("invalid category ID")), nil
	}

	category, err := store.GetCategoryByID(ctx, int32(categoryID))
	if err != nil {
		return utils.ErrorResponse(http.StatusInternalServerError, fmt.Errorf("error searching category by ID: %v", err)), nil
	}

	return utils.Response(http.StatusOK, category)
}
