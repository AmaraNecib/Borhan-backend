package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	database "github.com/AmaraNecib/Borhan-backend/Database"
	"github.com/AmaraNecib/Borhan-backend/api"
	"github.com/AmaraNecib/Borhan-backend/repository"
)

func main() {
	dbConn, err := database.ConnectToDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer dbConn.Close()

	// Create a new instance of the generated queries
	queries := repository.New(dbConn)
	if err := executeSQLFile(dbConn, "./sql/schema.sql"); err != nil {
		log.Fatalf("Failed to execute SQL file: %v", err)
	}
	fmt.Println("Successfully created schema")
	_, err = api.Init(queries)
	if err != nil {
		panic(err)
	} // Assuming you have generated the db package
	defer dbConn.Close()
}
func executeSQLFile(db *sql.DB, filePath string) error {
	// Read the SQL file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read file %s: %w", filePath, err)
	}

	// Execute the SQL commands in the file
	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("could not execute SQL: %w", err)
	}

	return nil
}
