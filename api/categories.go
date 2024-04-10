package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/otaviomart1ns/finsys/db/sqlc"
	"github.com/otaviomart1ns/finsys/utils"
)

type addCategoryRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) addCategory(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req addCategoryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	params := db.AddCategoryParams{
		UserID:      req.UserID,
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
	}

	category, err := server.store.AddCategory(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, category)
}

type updateCategoryRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (server *Server) updateCategory(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req updateCategoryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	params := db.UpdateCategoryParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
	}

	category, err := server.store.UpdateCategory(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, category)
}

type deleteCategoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteCategory(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req deleteCategoryRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.DeleteCategory(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, true)
}

type getCategoriesRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (server *Server) getCategories(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req getCategoriesRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := db.GetCategoriesParams{
		UserID:      req.UserID,
		Type:        req.Type,
		Title:       req.Title,
		Description: req.Description,
	}

	categories, err := server.store.GetCategories(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, categories)
}

type getCategoryByIDRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getCategoryByID(ctx *gin.Context) {
	errValiteToken := utils.GetTokenAndVerify(ctx)
	if errValiteToken != nil {
		return
	}

	var req getCategoryByIDRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	category, err := server.store.GetCategoryByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, category)
}
