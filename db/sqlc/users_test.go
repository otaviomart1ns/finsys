package db

import (
	"context"
	"testing"

	"github.com/otaviomart1ns/finsys/utils"
	"github.com/stretchr/testify/require"
)

func addRandomUser(t *testing.T) User {
	firstName, lastName := utils.RandomName()
	password := utils.RandomString(8)
	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	arg := AddUserParams{
		Username: utils.RandomString(6),
		Password: hashedPassword,
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

func TestUpdateUser(t *testing.T) {
	user := addRandomUser(t)

	newUsername := utils.RandomString(8)
	newPassword := utils.RandomString(10)
	newHashedPassword, err := utils.HashPassword(newPassword)
	require.NoError(t, err)
	newFirstName, newLastName := utils.RandomName()
	newBirth := utils.RandomBirthDate()
	newEmail := utils.RandomEmail(newFirstName, newLastName)

	tests := []struct {
		desc   string
		update UpdateUserParams
		check  func(*testing.T, User)
	}{
		{
			desc: "UpdateUsername",
			update: UpdateUserParams{
				ID:       user.ID,
				Username: newUsername,
				Password: user.Password,
				Name:     user.Name,
				LastName: user.LastName,
				Birth:    user.Birth,
				Email:    user.Email,
			},
			check: func(t *testing.T, updated User) {
				require.Equal(t, newUsername, updated.Username)
				require.Equal(t, user.Password, updated.Password)
				require.Equal(t, user.Name, updated.Name)
				require.Equal(t, user.LastName, updated.LastName)
				require.Equal(t, user.Birth, updated.Birth)
				require.Equal(t, user.Email, updated.Email)
			},
		},
		{
			desc: "UpdatePassword",
			update: UpdateUserParams{
				ID:       user.ID,
				Password: newHashedPassword,
				Username: user.Username,
				Name:     user.Name,
				LastName: user.LastName,
				Birth:    user.Birth,
				Email:    user.Email,
			},
			check: func(t *testing.T, updated User) {
				require.NotEqual(t, user.Password, updated.Password)
			},
		},
		{
			desc: "UpdateName",
			update: UpdateUserParams{
				ID:       user.ID,
				Name:     newFirstName,
				Username: user.Username,
				Password: user.Password,
				LastName: user.LastName,
				Birth:    user.Birth,
				Email:    user.Email,
			},
			check: func(t *testing.T, updated User) {
				require.Equal(t, newFirstName, updated.Name)
			},
		},
		{
			desc: "UpdateLastName",
			update: UpdateUserParams{
				ID:       user.ID,
				LastName: newLastName,
				Username: user.Username,
				Password: user.Password,
				Name:     user.Name,
				Birth:    user.Birth,
				Email:    user.Email,
			},
			check: func(t *testing.T, updated User) {
				require.Equal(t, newLastName, updated.LastName)
			},
		},
		{
			desc: "UpdateEmail",
			update: UpdateUserParams{
				ID:       user.ID,
				Email:    newEmail,
				Username: user.Username,
				Password: user.Password,
				Name:     user.Name,
				LastName: user.LastName,
				Birth:    user.Birth,
			},
			check: func(t *testing.T, updated User) {
				require.Equal(t, newEmail, updated.Email)
			},
		},
		{
			desc: "UpdateBirth",
			update: UpdateUserParams{
				ID:       user.ID,
				Birth:    newBirth,
				Username: user.Username,
				Password: user.Password,
				Name:     user.Name,
				LastName: user.LastName,
				Email:    user.Email,
			},
			check: func(t *testing.T, updated User) {
				require.Equal(t, newBirth.Format("2006-01-02"), updated.Birth.Format("2006-01-02"))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			updatedUser, err := testQueries.UpdateUser(context.Background(), tc.update)
			require.NoError(t, err)
			require.NotEmpty(t, updatedUser)

			tc.check(t, updatedUser)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	user := addRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
}
