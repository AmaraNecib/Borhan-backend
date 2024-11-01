package main

import (
	"log"

	database "github.com/AmaraNecib/Borhan-backend/Database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Database connection string
	// dbURL := "postgres://sensei57:aB3$c!@D#5f~G+?*@101.46.70.58:5432/postgres"

	// Connect to the database
	DATABASE, err := database.ConnectToDB()

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		database.CloseDB(DATABASE)
	}
	defer DATABASE.Close()
	// Example route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected to PostgreSQL!")
	})

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
