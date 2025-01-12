package api

import (
	"encoding/json"
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
		if err != nil {
			fmt.Println("here ", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": fmt.Sprintf("something went wrong: %v", err),
			})
		}
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
		id, err := db.GetClinicByEmail(c.Context(), Login.Email)
		if err != nil {
			fmt.Println("here ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Invalid Credentials",
			})
		}
		// Create token
		token, err := auth.CreateToken(id, "clinic")
		if err != nil {
			fmt.Println("here ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Invalid Credentials",
			})
		}

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"ok":     true,
			"token":  token,
			"userId": res.ID,
			// "token.id": auth.GetUserID(token),
			"role": "clinic"})
	}
}

func GetUserHistory(db *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		nationalID := c.Params("national_id")
		history, err := db.GetPatientHistoryByNationalId(c.Context(), nationalID)
		fmt.Println("history ", history[0].CreatedAt.Time)
		// get the gender from examination_data
		// gender :=
		//  add all ExaminationData and ExaminationsType from the history to data variable
		data := []types.GetPatientHistoryByNationalIdRow{}
		for _, v := range history {
			var examData types.HeartPredictRes
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"ok":  false,
					"msg": "Invalid Credentials",
				})
			}
			if err := json.Unmarshal(history[0].ExaminationData, &examData); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"ok":  false,
					"msg": "Error unmarshalling examination data",
				})
			}
			fmt.Println("examData ", examData)
			examData.CreatedAt = v.CreatedAt.Time.String()
			data = append(data, types.GetPatientHistoryByNationalIdRow{
				ExaminationData:  examData,
				ExaminationsType: v.ExaminationsType,
			})
		}
		var examData types.HeartPredictRes
		if err := json.Unmarshal(history[0].ExaminationData, &examData); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Error unmarshalling examination data",
			})
		}
		// reverses the data array
		for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":          true,
			"first Name":  history[0].FirstName,
			"last Name":   history[0].LastName,
			"national ID": history[0].NationalID,
			"Birth Date":  history[0].DateOfBirth,
			"gender":      examData.Sex,
			"examData":    data,
			// "history":     history[0].ExaminationsType,
		})
		// return c.Status(fiber.StatusOK).JSON(fiber.Map{
		// 	"ok":          true,
		// 	"first Name":  history[0].FirstName,
		// 	"last Name":   history[0].LastName,
		// 	"national ID": history[0].NationalID,
		// 	"Birth Date":  history[0].DateOfBirth,
		// 	"gender":      history[0].ExaminationData.Sex,
		// 	"history":     history[0].ExaminationsType,
		// })
	}
}

func GetAllPatient(db *repository.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		patients, err := db.GetAllPatients(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Invalid Credentials",
			})
		}
		// reverse the patients array
		for i, j := 0, len(patients)-1; i < j; i, j = i+1, j-1 {
			patients[i], patients[j] = patients[j], patients[i]
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":       true,
			"patients": patients,
		})
	}
}
