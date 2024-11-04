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
INSERT INTO Patients (national_id)
    VALUES ($1)  -- Replace with the actual national_id
    ON CONFLICT (national_id) DO NOTHING
    RETURNING id;

-- name: Info :many
SELECT current_database();

-- name: GetTables :many
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public';
