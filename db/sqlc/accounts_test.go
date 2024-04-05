package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/otaviomart1ns/finsys/utils"
	"github.com/stretchr/testify/require"
)

func addRandomAccount(t *testing.T) Account {
	category := addRandomCategory(t)

	params := AddAccountParams{
		UserID:      category.UserID,
		CategoryID:  category.ID,
		Title:       utils.RandomString(12),
		Type:        category.Type,
		Description: utils.RandomString(20),
		Value:       utils.RandomAccountValue(),
		Date:        time.Now(),
	}

	i, err := testQueries.AddAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, params.UserID, i.UserID)
	require.Equal(t, params.CategoryID, i.CategoryID)
	require.Equal(t, params.Value, i.Value)
	require.Equal(t, params.Title, i.Title)
	require.Equal(t, params.Type, i.Type)
	require.Equal(t, params.Description, i.Description)

	require.NotEmpty(t, i.CreatedAt)
	require.NotEmpty(t, i.Date)

	return i
}

func assertAccountsEquals(t *testing.T, expected, actual Account) {
	require.Equal(t, expected.UserID, actual.UserID)
	require.Equal(t, expected.CategoryID, actual.CategoryID)
	require.Equal(t, expected.Value, actual.Value)
	require.Equal(t, expected.Title, actual.Title)
	require.Equal(t, expected.Type, actual.Type)
	require.Equal(t, expected.Description, actual.Description)
	require.NotEmpty(t, actual.Date)
	require.NotEmpty(t, actual.CreatedAt)
}

func TestAddAccount(t *testing.T) {
	addRandomAccount(t)
}

func TestGetAccounts(t *testing.T) {
	account := addRandomAccount(t)

	params := GetAccountsParams{
		UserID: account.UserID,
		Type:   account.Type,
		CategoryID: sql.NullInt32{
			Valid: true,
			Int32: account.CategoryID,
		},
		Date: sql.NullTime{
			Valid: true,
			Time:  account.Date,
		},
		Title:       account.Title,
		Description: account.Description,
	}

	i, err := testQueries.GetAccounts(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	for _, a := range i {
		require.Equal(t, account.ID, a.ID)
		require.Equal(t, account.UserID, a.UserID)
		require.Equal(t, account.Title, a.Title)
		require.Equal(t, account.Description, a.Description)
		require.Equal(t, account.Value, a.Value)
		require.NotEmpty(t, a.CreatedAt)
		require.NotEmpty(t, a.Date)
	}
}

func TestGetAccountByID(t *testing.T) {
	account := addRandomAccount(t)

	i, err := testQueries.GetAccountByID(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	assertAccountsEquals(t, account, i)
}

func TestGetAccountByUser(t *testing.T) {
	account := addRandomAccount(t)

	i, err := testQueries.GetAccountByUser(context.Background(), account.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	assertAccountsEquals(t, account, i)
}

func TestGetAccountByCategory(t *testing.T) {
	account := addRandomAccount(t)

	i, err := testQueries.GetAccountByCategory(context.Background(), account.CategoryID)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	assertAccountsEquals(t, account, i)
}
func TestGetAccountReports(t *testing.T) {
	account := addRandomAccount(t)

	params := GetAccountReportsParams{
		UserID: account.UserID,
		Type:   account.Type,
	}

	i, err := testQueries.GetAccountReports(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)
}

func TestGetAccountGraph(t *testing.T) {
	account := addRandomAccount(t)

	params := GetAccountGraphParams{
		UserID: account.UserID,
		Type:   account.Type,
	}

	i, err := testQueries.GetAccountGraph(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)
}

func TestUpdateAccount(t *testing.T) {
	account := addRandomAccount(t)

	params := UpdateAccountParams{
		ID:          account.ID,
		Title:       utils.RandomString(12),
		Description: utils.RandomString(20),
		Value:       utils.RandomAccountValue(),
	}

	i, err := testQueries.UpdateAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, account.ID, i.ID)
	require.Equal(t, params.Title, i.Title)
	require.Equal(t, params.Description, i.Description)
	require.Equal(t, params.Value, i.Value)
	require.NotEmpty(t, i.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account := addRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}
