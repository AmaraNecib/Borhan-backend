package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func ConnectToDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	serviceURI := os.Getenv("DATABASE_URL")

	conn, _ := url.Parse(serviceURI)

	db, err := sql.Open("postgres", conn.String())
	if err != nil {
		fmt.Println("Error opening connection to database: ", err)
		return nil, err
	}

	return db, nil
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, err
	// }
	// serviceURI := os.Getenv("DATABASE_URL")

	// // conn, _ := url.Parse(serviceURI)

	// config, err := pgxpool.ParseConfig(serviceURI)
	// if err != nil {
	// 	return nil, err
	// }

	// db, err := pgxpool.NewWithConfig(context.Background(), config)
	// if err != nil {
	// 	return nil, err
	// }

	// return db, nil
}

func CloseDB(db *sql.DB) {
	db.Close()
}
