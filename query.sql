-- name: CreateUser :exec
INSERT INTO Users (email, password)
VALUES ($1, $2);