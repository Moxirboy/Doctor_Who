package rest

import (
	"html/template"
	"log"
	"net/http"
	_const "DoctorWho/internal/common/const"
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/delivery/html"
	"DoctorWho/internal/domain"
	"time"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (c controller) GetProgramForWeightLoss(ctx *gin.Context) {
	s := sessions.Default(ctx)
	var id = s.Get("userId").(int)
	var date string
	log.Println(s.Get("date"))
	dateInterface := s.Get("date")
	if dateInterface!=nil{
		date=dateInterface.(string)
	}else{
		date="null"
	}
	match := date == time.Now().Format("2006-01-02")
	if !match {
		_, _, _ = c.usecase.AutoExercise(id, dto.ProgramType(_const.StressWork))
		_, date, _ := c.usecase.AutoExercise(id, dto.ProgramType(_const.WeightLoss))
		s.Set("date", date)
		s.Save()
	}
	exercise, err := c.usecase.GetProgramForWeightLoss(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "internel server error",
		})
		return
	}
	ctx.JSON(200, exercise)
}

func (c controller) GetProgramForStress(ctx *gin.Context) {
	s := sessions.Default(ctx)
	var id = s.Get("userId").(int)
	var date string
	log.Println(s.Get("date"))
	dateInterface := s.Get("date")
	if dateInterface!=nil{
		date=dateInterface.(string)
	}else{
		date="null"
	}
	match := date == time.Now().Format("2006-01-02")
	if !match {
		_, _, _ = c.usecase.AutoExercise(id, dto.ProgramType(_const.StressWork))
		_, date, _ := c.usecase.AutoExercise(id, dto.ProgramType(_const.WeightLoss))
		s.Set("date", date)
		s.Save()
	}
	log.Println(id)
	exercise, err := c.usecase.GetProgramForStress(id)
	log.Println(exercise)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "internel server error",
		})
		return
	}
	ctx.JSON(200, exercise)
}
func (c controller) GetProgressStreesWork(ctx *gin.Context) {
	var personal domain.PersonalExercisesDone
	s := sessions.Default(ctx)
	personal.UserId = s.Get("userId").(int)
	if personal.UserId == 0 {
		ctx.JSON(200, gin.H{
			"progress ": 0,
		})
		return
	}
	var prog float64
	var err error
	if personal.UserId == 0 {
		prog = 0.0
	} else {
		personal.Typo = string(dto.StressWork)
		prog, err = c.usecase.GetProgress(personal)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "internel server error",
			})
			return
		}
	}

	ctx.JSON(200, gin.H{
		"progress ": prog,
	})
}
func (c controller) GetProgressWeightLoss(ctx *gin.Context) {
	var personal domain.PersonalExercisesDone
	var prog float64
	var err error
	s := sessions.Default(ctx)
	personal.UserId = s.Get("userId").(int)
	if personal.UserId == 0 {
		ctx.JSON(200, gin.H{
			"progress ": 0,
		})
		return
	}
	if personal.UserId == 0 {
		prog = 0.0
	} else {
		personal.Typo = string(dto.WeightLoss)
		prog, err = c.usecase.GetProgress(personal)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "internel server error",
			})
			return
		}
	}
	ctx.JSON(200, gin.H{
		"progress ": prog,
	})
}

func (cr controller) ProgramHandler(c *gin.Context) {
	tmpl, err := template.New("program").Parse(html.ProgramHtml)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	programs, err := cr.usecase.GetAllPrograms()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	programsJSON, err := json.Marshal(programs)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(c.Writer, domain.ProgramsPageData{ProgramsJSON: template.JS(programsJSON)})
}

func (cr controller) NewProgram(c *gin.Context) {

	var newProgram domain.Program
	newProgram.AgeUp = c.PostForm("AgeUp")
	log.Println(newProgram.AgeUp)
	if newProgram.AgeUp == "" {
		
		c.JSON(http.StatusBadRequest, gin.H{"error": "age up cannot be empty"})
		return
	}
	newProgram.AgeDown = c.PostForm("AgeDown")
	if newProgram.AgeDown == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AgeDown  cannot be empty"})
		return
	}
	newProgram.BMIUp = c.PostForm("BMIUp")
	if newProgram.BMIUp == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": " BMIUp cannot be empty"})
		return
	}
	newProgram.BMIDown = c.PostForm("BMIDown")
	if newProgram.AgeUp == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bmiDown cannot be empty"})
		return
	}
	newProgram.Type = c.PostForm("Type")
	if newProgram.AgeUp == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bmiDown cannot be empty"})
		return
	}
	_, err := cr.usecase.CreateProgram(newProgram)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Reload the page after form submission
	c.Status(200)
	c.Redirect(http.StatusSeeOther, "/save/program")
}
