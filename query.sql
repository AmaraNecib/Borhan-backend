-- name: CreateUser :one
INSERT INTO Users (email, password)
VALUES ($1, $2) RETURNING *;
-- name: CreateClinic :one
INSERT INTO clinics (clinic_name, user_id)
VALUES ($2, $1) RETURNING *;
-- name: GetUserByEmail :one
SELECT * FROM Users WHERE email = $1;

-- name: GetAllClinics :many
SELECT * FROM Clinics offset $1 limit $2;

-- name: CreatePatient :one
INSERT INTO Patients (last_name, first_name, date_of_birth, national_id)
VALUES ($1, $2, $3, $4)  -- Replace with the actual national_id
ON CONFLICT (national_id) 
DO UPDATE SET national_id = EXCLUDED.national_id  -- Dummy update to trigger the conflict handling
RETURNING id;


-- name: CreateExamination :one
INSERT INTO Examinations (patient_id, clinic_id, examinations_type, examination_data)
    VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetClinicByID :one
SELECT * FROM Clinics WHERE id = $1;

-- name: GetClinicByEmail :one
SELECT Clinics.id
FROM Clinics
JOIN Users ON Clinics.user_id = Users.id
WHERE Users.email = $1;  -- Replace $1 with the user email parameter
