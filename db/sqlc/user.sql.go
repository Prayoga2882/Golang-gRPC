// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
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
    fullname,
    email
) VALUES (
             $1,
             $2,
             $3,
             $4
         ) RETURNING username, hashed_password, fullname, email, password_change_at, created_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.Fullname,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.Fullname,
		&i.Email,
		&i.PasswordChangeAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, hashed_password, fullname, email, password_change_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.Fullname,
		&i.Email,
		&i.PasswordChangeAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE($1, hashed_password),
  password_change_at = COALESCE($2, password_change_at),
  fullname = COALESCE($3, fullname),
  email = COALESCE($4, email)
WHERE
  username = $5
RETURNING username, hashed_password, fullname, email, password_change_at, created_at
`

type UpdateUserParams struct {
	HashedPassword   sql.NullString `json:"hashed_password"`
	PasswordChangeAt sql.NullTime   `json:"password_change_at"`
	Fullname         sql.NullString `json:"fullname"`
	Email            sql.NullString `json:"email"`
	Username         string         `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangeAt,
		arg.Fullname,
		arg.Email,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.Fullname,
		&i.Email,
		&i.PasswordChangeAt,
		&i.CreatedAt,
	)
	return i, err
}
