package db

import (
	"context"
	"testing"

	"github.com/otaviomart1ns/finsys/utils"
	"github.com/stretchr/testify/require"
)

func addRandomUser(t *testing.T) User {
	firstName, lastName := utils.RandomName()

	arg := AddUserParams{
		Username: utils.RandomString(6),
		Password: utils.RandomString(8),
		Name:     firstName,
		LastName: lastName,
		Birth:    utils.RandomBirthDate(),
		Email:    utils.RandomEmail(firstName, lastName),
	}

	i, err := testQueries.AddUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	return i
}

func assertUserEquals(t *testing.T, expected, actual User) {
	require.Equal(t, expected.Username, actual.Username)
	require.Equal(t, expected.Password, actual.Password)
	require.Equal(t, expected.Name, actual.Name)
	require.Equal(t, expected.LastName, actual.LastName)
	require.Equal(t, expected.Birth.Format("2006-01-02"), actual.Birth.Format("2006-01-02"))
	require.Equal(t, expected.Email, actual.Email)
	require.NotEmpty(t, actual.CreatedAt)
}

func TestAddUser(t *testing.T) {
	addRandomUser(t)
}

func TestGetUsers(t *testing.T) {
	i, err := testQueries.GetUsers(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, i)
}

func TestGetUserByID(t *testing.T) {
	user := addRandomUser(t)

	i, err := testQueries.GetUserByID(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	assertUserEquals(t, user, i)
}

func TestGetUserByUsername(t *testing.T) {
	user := addRandomUser(t)

	i, err := testQueries.GetUserByUsername(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	assertUserEquals(t, user, i)
}

func TestGetUserByEmailAndPassword(t *testing.T) {
	user := addRandomUser(t)

	params := GetUserByEmailAndPasswordParams{
		Email:    user.Email,
		Password: user.Password,
	}

	i, err := testQueries.GetUserByEmailAndPassword(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	assertUserEquals(t, user, i)
}

func TestUpdateUser(t *testing.T) {
	user := addRandomUser(t)

	firstName, lastName := utils.RandomName()

	params := UpdateUserParams{
		ID:       user.ID,
		Username: utils.RandomString(8),
		Password: utils.RandomString(10),
		Name:     firstName,
		LastName: lastName,
		Birth:    utils.RandomBirthDate(),
		Email:    utils.RandomEmail(firstName, lastName),
	}

	i, err := testQueries.UpdateUser(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, params.ID, i.ID)
	require.Equal(t, params.Username, i.Username)
	require.Equal(t, params.Password, i.Password)
	require.Equal(t, params.Name, i.Name)
	require.Equal(t, params.LastName, i.LastName)
	require.Equal(t, params.Birth.Format("2006-01-02"), i.Birth.Format("2006-01-02"))
	require.Equal(t, params.Email, i.Email)
	require.NotEmpty(t, i.CreatedAt)
}

func TestDeleteUser(t *testing.T) {
	user := addRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
}
