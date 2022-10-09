-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email,
  url_count,
  user_type
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  email = COALESCE(sqlc.narg(email), email),
  url_count = COALESCE(sqlc.narg(url_count), url_count)
WHERE
  username = sqlc.arg(username)
RETURNING *;

-- name: IncUrlCount :one
UPDATE users
SET url_count = url_count + 1
WHERE username = sqlc.arg(username)
RETURNING *;

-- name: DecUrlCount :one
UPDATE users
SET url_count = url_count - 1
WHERE username = sqlc.arg(username)
RETURNING *;