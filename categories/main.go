package main

import (
	"context"
	"encoding/json"
	"fmt"
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

type getCategoriesRequest struct {
	UserID      int32  `form:"user_id" json:"user_id" binding:"required"`
	Type        string `form:"type" json:"type" binding:"required"`
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
}

func GetCategories(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	_, err := utils.GetTokenAndVerify(request)
	if err != nil {
		return utils.ErrorResponse(http.StatusUnauthorized, err), nil
	}

	var req getCategoriesRequest
	err = json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return utils.ErrorResponse(http.StatusBadRequest, err), nil
	}

	params := commonDB.GetCategoriesParams{
		UserID:      req.UserID,
		Type:        req.Type,
		Title:       req.Title,
		Description: req.Description,
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
