package domain
import "html/template"
type Exercise struct{
	Id string `json:"id"`
	Name string `json:"name"`
	ProgramId int `json:"program_id"`
	Info string `json:"info"`
	Link string `json:"link"`
		
	}
	type MedicalProgramsPageData struct {
		MedicalProgramsJSON template.JS
	}
	