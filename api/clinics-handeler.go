package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AmaraNecib/Borhan-backend/helpers"
	auth "github.com/AmaraNecib/Borhan-backend/jwt"
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
			fmt.Println("here")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": err,
			})
		}
		dateOfBirth, err := time.Parse("02/01/2006", heart.DateOfBirth)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": "the date of birth should be like this 15/05/1985",
			})
		}
		patient, err := db.CreatePatient(c.Context(), repository.CreatePatientParams{
			FirstName:   heart.FirstName,
			LastName:    heart.LastName,
			NationalID:  heart.NationalId,
			DateOfBirth: dateOfBirth,
		})
		if err != nil {
			fmt.Println("here On create patient")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": err,
			})
		}
		// culc the age from today - date of birth and convert it to the years as int
		age := int(time.Since(dateOfBirth).Hours() / 24 / 365)
		fmt.Println("patient ", patient)
		var gender string
		if heart.Sex == 1 {
			gender = "Male"
		} else {
			gender = "Female"
		}
		// check if the data is valid
		if !helpers.IsValidCpType(heart.Cp) || !helpers.IsValidCaType(heart.Ca) || !helpers.IsValidThalType(heart.Thal) || !helpers.IsValidLogic(heart.Fbs) || !helpers.IsValidLogic(heart.Exang) || !helpers.IsValidRestecgType(heart.Restecg) || !helpers.IsValidSlopeType(heart.Slope) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":      false,
				"msg":     "Invalid data",
				"heart":   helpers.IsValidCpType(heart.Cp),
				"ca":      helpers.IsValidCaType(heart.Ca),
				"thal":    helpers.IsValidThalType(heart.Thal),
				"fbs":     helpers.IsValidLogic(heart.Fbs),
				"exang":   helpers.IsValidLogic(heart.Exang),
				"restecg": helpers.IsValidRestecgType(heart.Restecg),
				"slope":   helpers.IsValidSlopeType(heart.Slope),
			})
		}

		heartInfo := types.HeartPredictInfo{
			Age:      age,
			Sex:      gender,
			Cp:       heart.Cp,
			Trestbps: heart.Trestbps,
			Chol:     heart.Chol,
			Fbs:      heart.Fbs,
			Restecg:  heart.Restecg,
			Thalach:  heart.Thalach,
			Exang:    heart.Exang,
			Oldpeak:  heart.Oldpeak,
			Slope:    heart.Slope,
			Ca:       heart.Ca,
			Thal:     heart.Thal,
		}
		heartInfoBytes, _ := json.Marshal(heartInfo)
		// fmt.Println(string(heartInfoBytes))
		// fmt.Printf("%v %v %v %v\n", patient, auth.GetUserID(strings.Split(c.Get("Authorization"), " ")[1]), "heart", string(heartInfoBytes))
		clinicID := auth.GetUserID(strings.Split(c.Get("Authorization"), " ")[1])

		// Check if the clinic exists
		// _, err = db.GetClinicByID(c.Context(), clinicID) // Function to fetch clinic by ID
		// if err != nil {
		// 	// If clinic doesn't exist, return an error
		// 	fmt.Println("Clinic ID does not exist")
		// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"ok":  false,
		// 		"msg": "Clinic ID not found in the database.",
		// 	})
		// }

		// If clinic exists, proceed with creating the examination
		_, err = db.CreateExamination(c.Context(), repository.CreateExaminationParams{
			PatientID:        patient,
			ClinicID:         clinicID, // clinic_id should now be valid
			ExaminationsType: "heart",
			ExaminationData:  heartInfoBytes,
		})
		if err != nil {
			fmt.Println("here On create examination")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":  false,
				"msg": err,
			})
		}
		// fmt.Println(heartPredict)
		// send a post request to the model to get the prediction 101.46.66.167/predict with the data
		/*
					{
			    "age": 90,
			    "trestbps": 120,
			    "chol": 229,
			    "thalch": 150,
			    "oldpeak": 20,
			    "sex": 1,
			    "cp": "typical angina",
			    "fbs": false,
			    "restecg": "lv hypertrophy",
			    "exang": true,
			    "slope": "flat",
			    "ca": 2,
			    "thal": "reversable"
			}
		*/
		if heart.Restecg == "hypertrophy" {
			heart.Restecg = "lv hypertrophy"
			fmt.Println("restecg hypertrophy", heart.Restecg)
		}
		requestBody := map[string]interface{}{
			"age":      age,
			"trestbps": heart.Trestbps,
			"chol":     heart.Chol,
			"thalch":   heart.Thalach,
			"oldpeak":  heart.Oldpeak,
			"sex":      heart.Sex,
			"cp":       heart.Cp,
			"fbs":      heart.Fbs,
			"restecg":  heart.Restecg,
			"exang":    heart.Exang,
			"slope":    heart.Slope,
			"ca":       heart.Ca,
			"thal":     heart.Thal,
		}

		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Failed to prepare request",
			})
		}

		resp, err := http.Post("http://101.46.67.80:5000/predict", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Failed to get prediction",
			})
		}
		defer resp.Body.Close()
		// print the response
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Failed to read response",
			})
		}
		fmt.Println("response body:", string(body))
		// Parse the response body into a map
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Failed to parse prediction response",
			})
		}

		// Extract the prediction value
		prediction, ok := result["prediction"]
		// get the first element of the prediction
		prediction = prediction.([]interface{})[0]
		fmt.Println("prediction:", prediction)

		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":  false,
				"msg": "Invalid prediction format from model",
			})
		}

		// Convert prediction to integer percentage
		predict := int(prediction.(float64))
		// generate random number for the prediction
		// rand.Seed(time.Now().UnixNano())
		// predict := rand.Intn(100)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":     true,
			"msg":    "Heart prediction",
			"result": predict,
			// "patient": patient,
			// "req":     heart,
			// "predict": heartPredict,
		})
	}
}
