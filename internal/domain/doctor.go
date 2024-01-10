package domain
type Doctor struct{
	Id int `json:"id"`
	User_id int `json:"user_id"`
	Name string `json:"name"`
	Workplace string `json:"workplace"`
	Expirience string `json:"expirience"`
	Education string `json:"education"`
	Card string `json:"card"`
	Photo string `json:"photo"`
}
