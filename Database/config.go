package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/joho/godotenv"
)

func ConnectToDB() (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	serviceURI := os.Getenv("DATABASE_URL")

	// conn, _ := url.Parse(serviceURI)

	config, err := pgxpool.ParseConfig(serviceURI)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *pgxpool.Pool) {
	db.Close()
}
