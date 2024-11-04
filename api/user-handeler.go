package api

import (
	"fmt"

	auth "github.com/AmaraNecib/Borhan-backend/jwt"
	"github.com/AmaraNecib/Borhan-backend/repository"
	"github.com/AmaraNecib/Borhan-backend/types"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateClinic(db *repository.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var clinic types.CreateClinic
		if err := ctx.BodyParser(&clinic); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": err,
			})
		}
		if clinic.Email == "" || clinic.Password == "" || clinic.ClinicName == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": "All fields are required",
			})
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(clinic.Password), bcrypt.DefaultCost)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": "something went wrong",
			})
		}
		user, err := db.CreateUser(ctx.Context(), repository.CreateUserParams{
			Email:    clinic.Email,
			Password: string(hash),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		_, err = db.CreateClinic(ctx.Context(), repository.CreateClinicParams{
			UserID:     user.ID,
			ClinicName: clinic.ClinicName,
		})
		response := fiber.Map{
			"ok":  true,
			"msg": "clinic created successfully",
		}
		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}

func loginForClinics(db *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Login types.Login
		if err := c.BodyParser(&Login); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": err,
			})
		}
		// Throws Unauthorized error
		// if email != "john" || pass != "doe" {
		// 	return c.SendStatus(fiber.StatusUnauthorized)
		// }
		// get user if by email
		// get user by email

		res, err := db.GetUserByEmail(c.Context(), Login.Email)

		// Create the Claims
		if err != nil {
			fmt.Println("heress ", err, Login)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Invalid Credentials"},
			)
		}
		err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(Login.Password))
		// claims := auth.Claims{
		// 	"id":   res.ID,
		// 	"role": res.RoleName,
		// 	"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
		// }
		if err != nil {
			fmt.Println("hello ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Invalid Credentials"},
			)
		}
		// Create token
		token, err := auth.CreateToken(res.ID, "clinic")
		if err != nil {
			fmt.Println("here ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Invalid Credentials",
			})
		}

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"ok":    true,
			"token": token,
			"role":  "clinic"})
	}
}
