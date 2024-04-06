// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package db

import (
	"context"
	"time"
)

const addUser = `-- name: AddUser :one
INSERT INTO users (
  username, 
  password, 
  name, 
  last_name, 
  birth, 
  email, 
  created_at, 
  updated_at
) 
VALUES 
  ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id, username, password, name, last_name, birth, email, created_at, updated_at
`

type AddUserParams struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Birth    time.Time `json:"birth"`
	Email    string    `json:"email"`
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, addUser,
		arg.Username,
		arg.Password,
		arg.Name,
		arg.LastName,
		arg.Birth,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.LastName,
		&i.Birth,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM 
  users 
WHERE 
  id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmailAndPassword = `-- name: GetUserByEmailAndPassword :one
SELECT 
  id, username, password, name, last_name, birth, email, created_at, updated_at 
FROM 
  users 
WHERE 
  email = $1 
  AND password = $2
`

type GetUserByEmailAndPasswordParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) GetUserByEmailAndPassword(ctx context.Context, arg GetUserByEmailAndPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmailAndPassword, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.LastName,
		&i.Birth,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT 
  id, username, password, name, last_name, birth, email, created_at, updated_at 
FROM 
  users 
WHERE 
  id = $1 
LIMIT 
  1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.LastName,
		&i.Birth,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT 
  id, username, password, name, last_name, birth, email, created_at, updated_at 
FROM 
  users 
WHERE 
  username = $1 
LIMIT 
  1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.LastName,
		&i.Birth,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT 
  id, username, password, name, last_name, birth, email, created_at, updated_at 
FROM 
  users 
ORDER BY 
  name, 
  last_name
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Name,
			&i.LastName,
			&i.Birth,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE 
  users 
SET 
  username = $2, 
  password = $3, 
  name = $4, 
  last_name = $5, 
  birth = $6, 
  email = $7, 
  updated_at = NOW() 
WHERE 
  id = $1 RETURNING id, username, password, name, last_name, birth, email, created_at, updated_at
`

type UpdateUserParams struct {
	ID       int32     `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Birth    time.Time `json:"birth"`
	Email    string    `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Name,
		arg.LastName,
		arg.Birth,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Name,
		&i.LastName,
		&i.Birth,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
