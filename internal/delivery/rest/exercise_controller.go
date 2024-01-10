package rest

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"testDeployment/internal/delivery/html"
	"testDeployment/internal/domain"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func (cr controller) ExerciseHandler(c *gin.Context) {
	tmpl, err := template.New("exercise").Parse(html.ExerciseHtml)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	Exercises, err := cr.usecase.GetAllExercise()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	medicalProgramsJSON, err := json.Marshal(Exercises)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(c.Writer, domain.MedicalProgramsPageData{MedicalProgramsJSON: template.JS(medicalProgramsJSON)})
}
func (cr controller) Newxercise(c *gin.Context) {

	var newExercise domain.Exercise
	programId:=c.PostForm("program_id")
	if programId==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "programId cannot be empty"})
		return
	}
	Name:=c.PostForm("name")
	if Name==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "name cannot be empty"})
		return
	}
	info:=c.PostForm("info")
	if info==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "info cannot be empty"})
		return
	}
	link:=c.PostForm("link_to_video")
	if link==""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "link cannot be empty"})
		return
	}
	intValue, err := strconv.Atoi(programId)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "conversion"})
		return
	}
	newExercise=domain.Exercise{
		ProgramId: intValue,
		Name: Name,
		Info: info,
		Link: link,
	}

	if err := c.ShouldBind(&newExercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := cr.usecase.CreateExercise(newExercise)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Reload the page after form submission
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/save/exercise/?id=%v", id))
}
