// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email,
  url_count,
  user_type
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING username, hashed_password, full_name, email, password_changed_at, url_count, user_type, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	UrlCount       int64  `json:"url_count"`
	UserType       string `json:"user_type"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.UrlCount,
		arg.UserType,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.UrlCount,
		&i.UserType,
		&i.CreatedAt,
	)
	return i, err
}

const decUrlCount = `-- name: DecUrlCount :one
UPDATE users
SET url_count = url_count - 1
WHERE username = $1
RETURNING username, hashed_password, full_name, email, password_changed_at, url_count, user_type, created_at
`

func (q *Queries) DecUrlCount(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, decUrlCount, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.UrlCount,
		&i.UserType,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, full_name, email, password_changed_at, url_count, user_type, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.UrlCount,
		&i.UserType,
		&i.CreatedAt,
	)
	return i, err
}

const incUrlCount = `-- name: IncUrlCount :one
UPDATE users
SET url_count = url_count + 1
WHERE username = $1
RETURNING username, hashed_password, full_name, email, password_changed_at, url_count, user_type, created_at
`

func (q *Queries) IncUrlCount(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, incUrlCount, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.UrlCount,
		&i.UserType,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE($1, hashed_password),
  password_changed_at = COALESCE($2, password_changed_at),
  full_name = COALESCE($3, full_name),
  email = COALESCE($4, email),
  url_count = COALESCE($5, url_count)
WHERE
  username = $6
RETURNING username, hashed_password, full_name, email, password_changed_at, url_count, user_type, created_at
`

type UpdateUserParams struct {
	HashedPassword    sql.NullString `json:"hashed_password"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	FullName          sql.NullString `json:"full_name"`
	Email             sql.NullString `json:"email"`
	UrlCount          sql.NullInt64  `json:"url_count"`
	Username          string         `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.FullName,
		arg.Email,
		arg.UrlCount,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.UrlCount,
		&i.UserType,
		&i.CreatedAt,
	)
	return i, err
}
