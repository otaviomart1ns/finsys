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
		Password: utils.RandomString(10),
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

/*func TestGetUsers(t *testing.T) {
	user, err := testQueries.GetUsers(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, user)
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

func TestUpdateUser(t *testing.T) {
	user := addRandomUser(t)

	firstName, lastName := utils.RandomName()

	params := UpdateUserParams{
		ID:       user.ID,
		Username: utils.RandomString(8),
		Password: utils.RandomString(12),
		Name:     firstName,
		LastName: lastName,
		Birth:    utils.RandomBirthDate(),
		Email:    utils.RandomEmail(firstName, lastName),
	}

	err := testQueries.UpdateUser(context.Background(), params)
	require.NoError(t, err)
}

func TestUpdateUserByUsername(t *testing.T) {
	user := addRandomUser(t)

	params := UpdateUserByUsernameParams{
		ID:       user.ID,
		Username: utils.RandomString(8),
	}

	err := testQueries.UpdateUserByUsername(context.Background(), params)
	require.NoError(t, err)
}

func TestUpdateUserByPassword(t *testing.T) {
	user := addRandomUser(t)

	params := UpdateUserByPasswordParams{
		ID:       user.ID,
		Password: utils.RandomString(12),
	}

	err := testQueries.UpdateUserByPassword(context.Background(), params)
	require.NoError(t, err)
}

func TestUpdateUserByName(t *testing.T) {
	user := addRandomUser(t)

	firstName, _ := utils.RandomName()

	params := UpdateUserByNameParams{
		ID:   user.ID,
		Name: firstName,
	}

	err := testQueries.UpdateUserByName(context.Background(), params)
	require.NoError(t, err)
}

func TestUpdateUserByLastName(t *testing.T) {
	user := addRandomUser(t)

	_, lastName := utils.RandomName()

	params := UpdateUserByLastNameParams{
		ID:       user.ID,
		LastName: lastName,
	}

	err := testQueries.UpdateUserByLastName(context.Background(), params)
	require.NoError(t, err)
}

func TestUpdateUserByBirth(t *testing.T) {
	user := addRandomUser(t)

	params := UpdateUserByBirthParams{
		ID:    user.ID,
		Birth: utils.RandomBirthDate(),
	}

	err := testQueries.UpdateUserByBirth(context.Background(), params)
	require.NoError(t, err)
}

func TestUpdateUserByEmail(t *testing.T) {
	user := addRandomUser(t)

	firstName, lastName := utils.RandomName()

	params := UpdateUserByEmailParams{
		ID:    user.ID,
		Email: utils.RandomEmail(firstName, lastName),
	}

	err := testQueries.UpdateUserByEmail(context.Background(), params)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	user := addRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	//Se a primeira execução de "DeleteUser" foi bem sucedida,a segunda deve retornar nil
	err2 := testQueries.DeleteUser(context.Background(), user.ID)
	require.Nil(t, err2)
}
*/
