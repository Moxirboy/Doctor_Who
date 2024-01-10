package domain
type PersonalExercises struct{
	Id int `json:"exercise_id"`
	Name string `json:"name"`
	Info string `json:"info"`
	Done bool `json:"done"`
	Link string `json:"link"`
}
type MarkAsDone struct{
	Id int `json:"exercise_id"`
	Done bool `json:"done"`
}
type PersonalExercisesDone struct{
	UserId int 
	Typo string 
	Date string
}
