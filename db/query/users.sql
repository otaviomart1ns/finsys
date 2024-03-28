-- name: AddUser :one
INSERT INTO users (
  username,
  password,
  name,
  last_name,
  birth,
  email,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, NOW(), NOW()
) RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUserByNameAndLastName :one
SELECT *
FROM users
WHERE name = $1 AND last_name = $2;

-- name: GetUserByEmailAndPassword :one
SELECT *
FROM users
WHERE email = $1 AND password = $2;

-- name: GetUsers :many
SELECT *
FROM users
ORDER BY name, last_name;

-- name: UpdateUserUsername :exec
UPDATE users
SET username = $2, updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $2, updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserName :exec
UPDATE users
SET name = $2, updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserLastName :exec
UPDATE users
SET last_name = $2, updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserBirth :exec
UPDATE users
SET birth = $2, updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserEmail :exec
UPDATE users
SET email = $2, updated_at = NOW()
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET 
  username = $2, 
  password = $3, 
  name = $4, 
  last_name = $5, 
  birth = $6, 
  email = $7, 
  updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;