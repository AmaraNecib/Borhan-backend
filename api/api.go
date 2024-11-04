package api

import (
	"log"
	"os"
	"strings"

	auth "github.com/AmaraNecib/Borhan-backend/jwt"
	"github.com/AmaraNecib/Borhan-backend/repository"
	static "github.com/AmaraNecib/Borhan-backend/static"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	})
}
func Init(db *repository.Queries) (*fiber.App, error) {
	app := fiber.New(
		fiber.Config{
			Prefork: true,
		},
	)
	// app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		// AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger.New())
	// app.Get("/test", PredictHeartV1(db))

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected")
	})
	// Clinic
	clinic := v1.Group("/clinic")
	clinic.Post("/register", CreateClinic(db))
	clinic.Post("/login", loginForClinics(db))
	clinic.Get("/all/clinics", AdminsOnly, getAllClinics(db))
	// User
	// v1.Post("/register", CreateUser(db))
	// v1.Post("/login", login(db))
	// authorized routes
	v1.Get(("/token"), restricted)
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok":  false,
				"msg": "Unauthorized",
			})
		},
	}))
	clinic.Post("/predict/heart", PredictHeart(db))
	log.Fatal(app.Listen(":3000"))
	return app, nil
}

func restricted(c *fiber.Ctx) error {
	if auth.ValidToken(strings.Split(c.Get("Authorization"), " ")[1]) {
		return c.SendString("Welcome to the restricted area")
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"ok":  false,
			"msg": "Unauthorized",
		})
	}
}

func AdminsOnly(c *fiber.Ctx) error {
	if auth.GetUserRole(strings.Split(c.Get("Authorization"), " ")[1]) == static.Admin {
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"ok":  false,
			"msg": "this route does not exist",
		})
	}
}
