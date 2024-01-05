package dto

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}
type UserInfo struct {
	Name   string `json:"name"`
	Weigh  string `json:"weigh"`
	Height string `json:"height"`
	Age    string `json:"age"`
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
