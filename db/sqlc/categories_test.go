package db

import (
	"context"
	"testing"

	"github.com/otaviomart1ns/finsys/utils"
	"github.com/stretchr/testify/require"
)

func addRandomCategory(t *testing.T) Category {
	user := addRandomUser(t)

	params := AddCategoryParams{
		UserID:      user.ID,
		Title:       utils.RandomString(12),
		Type:        utils.RandomCategoryType(),
		Description: utils.RandomString(20),
	}

	i, err := testQueries.AddCategory(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, params.UserID, i.UserID)
	require.Equal(t, params.Title, i.Title)
	require.Equal(t, params.Type, i.Type)
	require.Equal(t, params.Description, i.Description)
	require.NotEmpty(t, i.CreatedAt)

	return i
}

func TestAddCategory(t *testing.T) {
	addRandomCategory(t)
}

func TestGetCategories(t *testing.T) {
	category := addRandomCategory(t)

	params := GetCategoriesParams{
		UserID:      category.UserID,
		Type:        category.Type,
		Title:       category.Title,
		Description: category.Description,
	}

	i, err := testQueries.GetCategories(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	for _, c := range i {
		require.Equal(t, category.ID, c.ID)
		require.Equal(t, category.UserID, c.UserID)
		require.Equal(t, category.Title, c.Title)
		require.Equal(t, category.Description, c.Description)
		require.NotEmpty(t, c.CreatedAt)
	}
}

func TestGetCategoryByID(t *testing.T) {
	category := addRandomCategory(t)

	i, err := testQueries.GetCategoryByID(context.Background(), category.ID)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, i.ID, category.ID)
	require.Equal(t, i.Title, category.Title)
	require.Equal(t, i.Description, category.Description)
	require.Equal(t, i.Type, category.Type)
	require.NotEmpty(t, category.CreatedAt)
}

func TestUpdateCategory(t *testing.T) {
	category := addRandomCategory(t)

	params := UpdateCategoryParams{
		ID:          category.ID,
		Title:       utils.RandomString(12),
		Description: utils.RandomString(20),
	}

	i, err := testQueries.UpdateCategory(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, category.ID, i.ID)
	require.Equal(t, params.Title, i.Title)
	require.Equal(t, params.Description, i.Description)
	require.NotEmpty(t, i.CreatedAt)
}

func TestDeleteCategory(t *testing.T) {
	category := addRandomCategory(t)

	err := testQueries.DeleteCategory(context.Background(), category.ID)
	require.NoError(t, err)

	//Se a primeira execução de "DeleteCategory" foi bem sucedida,a segunda deve retornar nil
	err2 := testQueries.DeleteCategory(context.Background(), category.ID)
	require.Nil(t, err2)
}
