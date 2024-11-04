package types

type Heart struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	NationalId  string `json:"national_id"`
	DateOfFirth string `json:"date_of_birth"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateClinic struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	ClinicName string `json:"clinic_name"`
}

type HeartPredictInfo struct {
	Age      int     `json:"age"`
	Sex      int     `json:"sex"`
	Cp       int     `json:"cp"`
	Trestbps float64 `json:"trestbps"`
	Chol     float64 `json:"chol"`
	Fbs      int     `json:"fbs"`
	Restecg  int     `json:"restecg"`
	Thalach  float64 `json:"thalach"`
	Exang    int     `json:"exang"`
	Oldpeak  float64 `json:"oldpeak"`
	Slope    int     `json:"slope"`
	Ca       int     `json:"ca"`
	Thal     int     `json:"thal"`
}

type HeartPredictReq struct {
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	NationalId  string  `json:"national_id"`
	DateOfBirth string  `json:"date_of_birth"`
	Sex         int     `json:"sex"`
	Cp          int     `json:"cp"`
	Trestbps    float64 `json:"trestbps"`
	Chol        float64 `json:"chol"`
	Fbs         int     `json:"fbs"`
	Restecg     int     `json:"restecg"`
	Thalach     float64 `json:"thalach"`
	Exang       int     `json:"exang"`
	Oldpeak     float64 `json:"oldpeak"`
	Slope       int     `json:"slope"`
	Ca          int     `json:"ca"`
	Thal        int     `json:"thal"`
}
