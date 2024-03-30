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

func TestGetUserByEmail(t *testing.T) {
	user := addRandomUser(t)

	i, err := testQueries.GetUserByEmail(context.Background(), user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	assertUserEquals(t, user, i)
}

func TestGetUserByNameAndLastName(t *testing.T) {
	user := addRandomUser(t)

	params := GetUserByNameAndLastNameParams{
		Name:     user.Name,
		LastName: user.LastName,
	}

	i, err := testQueries.GetUserByNameAndLastName(context.Background(), params)
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

	require.Equal(t, params.Username, i.Username)
	require.Equal(t, params.Password, i.Password)
	require.Equal(t, params.Name, i.Name)
	require.Equal(t, params.LastName, i.LastName)
	require.Equal(t, params.Birth.Format("2006-01-02"), i.Birth.Format("2006-01-02"))
	require.Equal(t, params.Email, i.Email)
	require.NotEmpty(t, i.CreatedAt)
}

func TestUpdateUserByUsername(t *testing.T) {
	user := addRandomUser(t)

	params := UpdateUserByUsernameParams{
		ID:       user.ID,
		Username: utils.RandomString(10),
	}

	i, err := testQueries.UpdateUserByUsername(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, params.Username, i.Username)
	require.NotEmpty(t, i.CreatedAt)
}

func TestUpdateUserByPassword(t *testing.T) {
	user := addRandomUser(t)

	params := UpdateUserByPasswordParams{
		ID:       user.ID,
		Password: utils.RandomString(12),
	}

	i, err := testQueries.UpdateUserByPassword(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, params.Password, i.Password)
	require.NotEmpty(t, i.CreatedAt)
}

func TestUpdateUserByName(t *testing.T) {
	user := addRandomUser(t)

	firstName, _ := utils.RandomName()

	params := UpdateUserByNameParams{
		ID:   user.ID,
		Name: firstName,
	}

	i, err := testQueries.UpdateUserByName(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, params.Name, i.Name)
	require.NotEmpty(t, i.CreatedAt)
}

func TestUpdateUserByLastName(t *testing.T) {
	user := addRandomUser(t)

	_, lastName := utils.RandomName()

	params := UpdateUserByLastNameParams{
		ID:       user.ID,
		LastName: lastName,
	}

	i, err := testQueries.UpdateUserByLastName(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, params.LastName, i.LastName)
	require.NotEmpty(t, i.CreatedAt)
}

func TestUpdateUserByBirth(t *testing.T) {
	user := addRandomUser(t)

	params := UpdateUserByBirthParams{
		ID:    user.ID,
		Birth: utils.RandomBirthDate(),
	}

	i, err := testQueries.UpdateUserByBirth(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, params.Birth.Format("2006-01-02"), i.Birth.Format("2006-01-02"))
	require.NotEmpty(t, i.CreatedAt)
}

func TestUpdateUserByEmail(t *testing.T) {
	user := addRandomUser(t)

	firstName, lastName := utils.RandomName()

	params := UpdateUserByEmailParams{
		ID:    user.ID,
		Email: utils.RandomEmail(firstName, lastName),
	}

	i, err := testQueries.UpdateUserByEmail(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, i)

	require.Equal(t, params.Email, i.Email)
	require.NotEmpty(t, i.CreatedAt)
}

func TestDeleteUser(t *testing.T) {
	user := addRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	//Se a primeira execução de "DeleteUser" foi bem sucedida,a segunda deve retornar nil
	err2 := testQueries.DeleteUser(context.Background(), user.ID)
	require.Nil(t, err2)
}
