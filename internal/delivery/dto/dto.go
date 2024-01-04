package dto

import "DoctorWho/internal/domain"

type User struct {
	Phone_number string `json:"phone_Number"`
}
type UserInfo struct {
	Name   string `json:"name"`
	Weigh  string `json:"weigh"`
	Height string `json:"height"`
	Age    string `json:"age"`
	Waist  string `json:"waist"`
}
type Program struct {
	Id          int                `json:"id"`
	ProgramType domain.ProgramType `json:"programType"`
}
