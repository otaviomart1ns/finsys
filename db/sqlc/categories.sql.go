// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: categories.sql

package db

import (
	"context"
)

const addCategory = `-- name: AddCategory :one
INSERT INTO categories (
  user_id, 
  title, 
  type, 
  description, 
  created_at, 
  updated_at
)
VALUES (
  $1, $2, $3, $4, NOW(), NOW()
)
RETURNING id, user_id, title, type, description, created_at, updated_at
`

type AddCategoryParams struct {
	UserID      int32  `json:"user_id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

func (q *Queries) AddCategory(ctx context.Context, arg AddCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, addCategory,
		arg.UserID,
		arg.Title,
		arg.Type,
		arg.Description,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategories = `-- name: GetCategories :many
SELECT id, user_id, title, type, description, created_at, updated_at 
FROM categories
WHERE user_id = $1
AND type = $2
AND
  LOWER(title) LIKE CONCAT('%', LOWER($3::text), '%')
AND
  LOWER(description) LIKE CONCAT('%', LOWER($4::text), '%')
`

type GetCategoriesParams struct {
	UserID      int32  `json:"user_id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getCategories,
		arg.UserID,
		arg.Type,
		arg.Title,
		arg.Description,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
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

const getCategoryByID = `-- name: GetCategoryByID :one
SELECT id, user_id, title, type, description, created_at, updated_at
FROM categories
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCategoryByID(ctx context.Context, id int32) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategoryByID, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories
SET
  title = $2,
  description = $3,
  updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, title, type, description, created_at, updated_at
`

type UpdateCategoryParams struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.ID, arg.Title, arg.Description)
	return err
}
