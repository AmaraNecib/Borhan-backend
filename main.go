package main

import (
	"log"

	"github.com/AmaraNecib/Borhan-backend/DB"
	database "github.com/AmaraNecib/Borhan-backend/Database"
	"github.com/AmaraNecib/Borhan-backend/api"
)

func main() {
	DATABASE, err := database.ConnectToDB()

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		database.CloseDB(DATABASE)
	}
	// schema, err := ioutil.ReadFile("schema.sql")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Execute the schema creation
	// if _, err := DATABASE.Exec(string(schema)); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Successfully created schema")
	queries := DB.New(DATABASE)
	_, err = api.Init(queries)
	if err != nil {
		panic(err)
	}
	defer DATABASE.Close()
}
