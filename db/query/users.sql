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
) 
VALUES 
  ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING *;

-- name: GetUserByID :one
SELECT 
  * 
FROM 
  users 
WHERE 
  id = $1 
LIMIT 
  1;

-- name: GetUserByUsername :one
SELECT 
  * 
FROM 
  users 
WHERE 
  username = $1 
LIMIT 
  1;

-- name: GetUserByEmailAndPassword :one
SELECT 
  * 
FROM 
  users 
WHERE 
  email = $1 
  AND password = $2;

-- name: GetUsers :many
SELECT 
  * 
FROM 
  users 
ORDER BY 
  name, 
  last_name;

-- name: UpdateUser :one
UPDATE 
  users 
SET 
  username = COALESCE($2, username),
  password = COALESCE($3, password),
  name = COALESCE($4, name),
  last_name = COALESCE($5, last_name),
  birth = COALESCE($6, birth),
  email = COALESCE($7, email),
  updated_at = NOW()
WHERE 
  id = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM 
  users 
WHERE 
  id = $1;
