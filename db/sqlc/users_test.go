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
		Username: utils.RandomUsernamePassword(6),
		Password: utils.RandomUsernamePassword(10),
		Name:     firstName,
		LastName: lastName,
		Birth:    utils.RandomBirthDate(),
		Email:    utils.RandomEmail(firstName, lastName),
	}

	user, err := testQueries.AddUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
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

func TestGetUserByID(t *testing.T) {
	arg := addRandomUser(t)

	user, err := testQueries.GetUserByID(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	assertUserEquals(t, arg, user)
}

func TestGetUserByUsername(t *testing.T) {
	arg := addRandomUser(t)

	user, err := testQueries.GetUserByUsername(context.Background(), arg.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	assertUserEquals(t, arg, user)
}

func TestGetUserByEmail(t *testing.T) {
	arg := addRandomUser(t)

	user, err := testQueries.GetUserByEmail(context.Background(), arg.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	assertUserEquals(t, arg, user)
}

func TestGetUserByNameAndLastName(t *testing.T) {
	arg := addRandomUser(t)

	params := GetUserByNameAndLastNameParams{
		Name:     arg.Name,
		LastName: arg.LastName,
	}

	user, err := testQueries.GetUserByNameAndLastName(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	assertUserEquals(t, arg, user)
}

func TestGetUserByEmailAndPassword(t *testing.T) {
	arg := addRandomUser(t)

	params := GetUserByEmailAndPasswordParams{
		Email:    arg.Email,
		Password: arg.Password,
	}

	user, err := testQueries.GetUserByEmailAndPassword(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	assertUserEquals(t, arg, user)
}

func TestGetUsers(t *testing.T) {
	user, err := testQueries.GetUsers(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
