package dto

type User struct {
	
	Email string `json:"email"`
}
type UserInfo struct {
	Id 	   int 
	Name   string `json:"name"`
	Weigh  string `json:"weigh"`
	Height string `json:"height"`
	Age    int `json:"age"`
	Gender string `json:"gender"`
	Waist  string `json:"waist"`
}

type Program struct {
	Id          int         `json:"id"`
	ProgramType ProgramType `json:"programType"`
}
type ProgramType string
type ProType string

const (
	WeightLoss  = ProgramType("weight_loss")
	StressWork  = ProgramType("stress_work")
	Recommended = ProType("recommended")
	Personal    = ProType("personal")
)
