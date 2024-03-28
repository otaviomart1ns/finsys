// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: accounts.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const addAccount = `-- name: AddAccount :one
INSERT INTO accounts (
  user_id,
  category_id,
  title,
  type,
  description,
  value,
  date,
  created_at, 
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
) RETURNING id, user_id, category_id, title, type, description, value, date, created_at, updated_at
`

type AddAccountParams struct {
	UserID      int32     `json:"user_id"`
	CategoryID  int32     `json:"category_id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Value       string    `json:"value"`
	Date        time.Time `json:"date"`
}

func (q *Queries) AddAccount(ctx context.Context, arg AddAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, addAccount,
		arg.UserID,
		arg.CategoryID,
		arg.Title,
		arg.Type,
		arg.Description,
		arg.Value,
		arg.Date,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccountByID = `-- name: GetAccountByID :one
SELECT id, user_id, category_id, title, type, description, value, date, created_at, updated_at 
FROM accounts
WHERE id = $1 
LIMIT 1
`

func (q *Queries) GetAccountByID(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByID, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccounts = `-- name: GetAccounts :many
SELECT
  a.id,
  a.user_id,
  a.title,
  a.type,
  a.description,
  a.value,
  a.date,
  a.created_at,
  c.title as category_title
FROM
  accounts a
LEFT JOIN
  categories c ON c.id = a.category_id
WHERE
  a.user_id = $1
AND
  a.type = $2
AND
  LOWER(a.title) LIKE CONCAT('%', LOWER($3::text), '%')
AND
  LOWER(a.description) LIKE CONCAT('%', LOWER($4::text), '%')
AND
  a.category_id = COALESCE($5, a.category_id)
AND
  a.date = COALESCE($6, a.date)
`

type GetAccountsParams struct {
	UserID      int32         `json:"user_id"`
	Type        string        `json:"type"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	CategoryID  sql.NullInt32 `json:"category_id"`
	Date        sql.NullTime  `json:"date"`
}

type GetAccountsRow struct {
	ID            int32          `json:"id"`
	UserID        int32          `json:"user_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         string         `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccounts,
		arg.UserID,
		arg.Type,
		arg.Title,
		arg.Description,
		arg.CategoryID,
		arg.Date,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsRow{}
	for rows.Next() {
		var i GetAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
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

const getAccountsByCategory = `-- name: GetAccountsByCategory :many
SELECT id, user_id, category_id, title, type, description, value, date, created_at, updated_at
FROM accounts
WHERE category_id = $1
`

func (q *Queries) GetAccountsByCategory(ctx context.Context, categoryID int32) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByCategory, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
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

const getAccountsByUser = `-- name: GetAccountsByUser :many
SELECT id, user_id, category_id, title, type, description, value, date, created_at, updated_at
FROM accounts
WHERE user_id = $1
`

func (q *Queries) GetAccountsByUser(ctx context.Context, userID int32) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
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

const getAccountsForUserByType = `-- name: GetAccountsForUserByType :many
SELECT id, user_id, category_id, title, type, description, value, date, created_at, updated_at
FROM accounts
WHERE user_id = $1 AND type = $2
`

type GetAccountsForUserByTypeParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) GetAccountsForUserByType(ctx context.Context, arg GetAccountsForUserByTypeParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsForUserByType, arg.UserID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoryID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
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

const getAccountsGraph = `-- name: GetAccountsGraph :one
SELECT COUNT(*) 
FROM accounts
WHERE user_id = $1 
AND type = $2
`

type GetAccountsGraphParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) GetAccountsGraph(ctx context.Context, arg GetAccountsGraphParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountsGraph, arg.UserID, arg.Type)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAccountsReports = `-- name: GetAccountsReports :one
SELECT SUM(value) AS sum_value 
FROM accounts
WHERE user_id = $1 
AND type = $2
`

type GetAccountsReportsParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) GetAccountsReports(ctx context.Context, arg GetAccountsReportsParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountsReports, arg.UserID, arg.Type)
	var sum_value int64
	err := row.Scan(&sum_value)
	return sum_value, err
}

const updateAccount = `-- name: UpdateAccount :exec
UPDATE accounts
SET 
  title = $2, 
  description = $3, 
  value = $4,
  updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, category_id, title, type, description, value, date, created_at, updated_at
`

type UpdateAccountParams struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) error {
	_, err := q.db.ExecContext(ctx, updateAccount,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Value,
	)
	return err
}

const updateAccountDepositInto = `-- name: UpdateAccountDepositInto :exec
UPDATE accounts
SET
  value = value + $2,
  updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, category_id, title, type, description, value, date, created_at, updated_at
`

type UpdateAccountDepositIntoParams struct {
	ID    int32  `json:"id"`
	Value string `json:"value"`
}

func (q *Queries) UpdateAccountDepositInto(ctx context.Context, arg UpdateAccountDepositIntoParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountDepositInto, arg.ID, arg.Value)
	return err
}

const updateAccountWithdrawFrom = `-- name: UpdateAccountWithdrawFrom :exec
UPDATE accounts
SET
  value = value - $2,
  updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, category_id, title, type, description, value, date, created_at, updated_at
`

type UpdateAccountWithdrawFromParams struct {
	ID    int32  `json:"id"`
	Value string `json:"value"`
}

func (q *Queries) UpdateAccountWithdrawFrom(ctx context.Context, arg UpdateAccountWithdrawFromParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountWithdrawFrom, arg.ID, arg.Value)
	return err
}
