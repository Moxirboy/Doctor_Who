package dto

type User struct {
	
	Email string `json:"email"`
}
type UserInfo struct {
	Id 	   int 
	Name   string `json:"name"`
	Age    int `json:"age"`
	Gender string `json:"gender"`
}