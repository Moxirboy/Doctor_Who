package rest

import (
	"DoctorWho/internal/domain"

	"github.com/gin-gonic/gin"
)
func(c controller) CreateDoctor(ctx *gin.Context){
	var newUser domain.NewUser
	ctx.ShouldBindJSON(&newUser)
	id,err:=c.usecase.RegisterDoctor(&newUser)
	if err!=nil{
		if err != nil {
			c.bot.SendErrorNotification(err)
			ctx.JSON(303, gin.H{
				"message":"user is registered",
			})
			return
		}
	}
	ctx.JSON(200,gin.H{
		"id":id,
	})
}

// func ( c controller) FillDoctorInfo(ctx *gin.Context){
// 	var Doctor domain.Doctor
// 	ctx.ShouldBindJSON(&Doctor)
// 	id ,err:=
// }