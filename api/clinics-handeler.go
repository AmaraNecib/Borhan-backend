package api

import (
	"fmt"
	"strconv"

	"github.com/AmaraNecib/Borhan-backend/repository"
	"github.com/AmaraNecib/Borhan-backend/types"
	"github.com/gofiber/fiber/v2"
)

func getAllClinics(db *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pageStr := c.Query("page", "1")          // Default to "1" if not provided
		sizeAgeStr := c.Query("page_size", "25") // Default to "25" if not provided

		// Convert to integers
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": "Invalid page number",
			})
		}

		sizeAge, err := strconv.Atoi(sizeAgeStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": "Invalid sizeage number",
			})
		}
		clinics, err := db.GetAllClinics(c.Context(), repository.GetAllClinicsParams{
			Offset: int32(page - 1),
			Limit:  int32(sizeAge),
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": err,
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":  true,
			"msg": clinics,
		})
	}
}

func PredictHeart(db *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var heart types.HeartPredictReq
		if err := c.BodyParser(&heart); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": err,
			})
		}
		patient, err := db.CreatePatient(c.Context(), heart.NationalId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": err,
			})
		}
		fmt.Println(patient)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":  true,
			"msg": "Heart prediction",
		})
	}
}

func PredictHeartV1(db *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// patient, err := db.CreatePatient(c.Context(), "123456789012345678")
		test, err := db.Info(c.Context())
		fmt.Println(test)
		if err != nil {
			fmt.Errorf("%w", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": test,
			})
		}
		// fmt.Println(patient)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":      true,
			"msg":     "Heart prediction",
			"patient": test,
		})
	}
}
