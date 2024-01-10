package domain
import "html/template"
type ProgramWeightLoss struct {
	id int
}
type Program struct {
	ID      string     `json:"id"`
	AgeUp   string     `json:"ageUp"`
	AgeDown string     `json:"ageDown"`
	BMIUp   string `json:"bmiUp"`
	BMIDown string `json:"bmiDown"`
	Type    string  `json:"type"`
	
}
type ProgramsPageData struct {
	ProgramsJSON template.JS
}