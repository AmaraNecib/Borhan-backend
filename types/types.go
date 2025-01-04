package types

// type slope enum =["downsloping", "flat", "upsloping"]
// Define allowed values for SlopeType
type SlopeType string

type CaType float32

type ThalType string

type Logic bool

type RestecgType string
type CpType string

const (
	Downsloping SlopeType = "downsloping"
	Flat        SlopeType = "flat"
	Upsloping   SlopeType = "upsloping"

	Zero  CaType = 0
	One   CaType = 1
	Two   CaType = 2
	Three CaType = 3

	FixedDefect      ThalType = "fixed defect"
	NormalThal       ThalType = "normal"
	ReversableDefect ThalType = "reversable defect"
	Reversable       ThalType = "reversable"

	False Logic = false
	True  Logic = true

	Hypertrophy        RestecgType = "hypertrophy"
	NormalRestecg      RestecgType = "normal"
	STTWaveAbnormality RestecgType = "st-t wave abnormality"

	//['typical angina', 'asymptomatic', 'non-anginal', 'atypical angina']
	TypicalAngina  CpType = "typical angina"
	Asymptomatic   CpType = "asymptomatic"
	NonAnginal     CpType = "non-anginal"
	AtypicalAngina CpType = "atypical angina"
)

type HeartPredictReq struct {
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	NationalId  string      `json:"national_id"`
	DateOfBirth string      `json:"date_of_birth"`
	Oldpeak     float64     `json:"oldpeak"`
	Sex         int         `json:"sex"`
	Cp          CpType      `json:"cp"`
	Trestbps    float64     `json:"trestbps"`
	Chol        float64     `json:"chol"`
	Fbs         Logic       `json:"fbs"`     //done
	Restecg     RestecgType `json:"restecg"` //done
	Thalach     float64     `json:"thalach"`
	Exang       Logic       `json:"exang"`
	Slope       SlopeType   `json:"slope"` //done
	Ca          CaType      `json:"ca"`    //done
	Thal        ThalType    `json:"thal"`  //done
}
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
	Age      int         `json:"age"`
	Oldpeak  float64     `json:"oldpeak"`
	Sex      string      `json:"sex"`
	Cp       CpType      `json:"cp"`
	Trestbps float64     `json:"trestbps"`
	Chol     float64     `json:"chol"`
	Fbs      Logic       `json:"fbs"`     //done
	Restecg  RestecgType `json:"restecg"` //done
	Thalach  float64     `json:"thalach"`
	Exang    Logic       `json:"exang"`
	Slope    SlopeType   `json:"slope"` //done
	Ca       CaType      `json:"ca"`    //done
	Thal     ThalType    `json:"thal"`  //done
}
