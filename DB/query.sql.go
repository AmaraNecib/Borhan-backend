// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package DB

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO Users (email, password)
VALUES ($1, $2)
`

type CreateUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Email, arg.Password)
	return err
}
